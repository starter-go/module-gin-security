package vericode

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/starter-go/application"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/mails"
	"github.com/starter-go/security/auth"
	"github.com/starter-go/security/jwt"
	"github.com/starter-go/security/random"
	"github.com/starter-go/vlog"

	"github.com/starter-go/security-gin-gorm/components/services"
)

// VerificationServiceImpl ...
type VerificationServiceImpl struct {

	//starter:component
	_as func(services.VerificationService) //starter:as("#")

	Sender      mails.Service      //starter:inject("#")
	JWT         jwt.Service        //starter:inject("#")
	UUIDService random.UUIDService //starter:inject("#")

	SenderFromSMS   string //starter:inject("${security.verification-code-sender.sms.from}")
	SenderFromMail  string //starter:inject("${security.verification-code-sender.mail.from}")
	TokenMaxAgeInMS int64  //starter:inject("${security.verification-code-token.max-age.ms}")

	saltGen saltGeneratingService
}

func (inst *VerificationServiceImpl) _impl() services.VerificationService {
	return inst
}

// Life ...
func (inst *VerificationServiceImpl) Life() *application.Life {
	return &application.Life{
		OnStartPre: inst.onInit,
	}
}

func (inst *VerificationServiceImpl) onInit() error {
	return inst.saltGen.init()
}

// Verify ...
func (inst *VerificationServiceImpl) Verify(c context.Context, v *services.Verification) error {

	// get
	token, err := inst.JWT.GetDTO(c)
	if err != nil {
		token = &jwt.Token{}
	}

	data1 := data4jwt{}
	data1.loadProperties(token.Properties)

	t1 := data1.t1
	t2 := data1.t2
	sum1 := data1.sum.Bytes()
	uuid := data1.uuid
	now := lang.Now()

	data2 := data4sum{}
	data2.account = inst.getAccount(v)
	data2.mechanism = v.Mechanism
	data2.code = v.Code
	data2.salt = inst.makeSalt(t1)
	data2.t1 = t1
	data2.t2 = t2
	data2.uuid = uuid

	sum2 := data2.computeHash()
	ok := bytes.Equal(sum1, sum2)
	if ok && t1 <= now && now <= t2 {
		return nil
	}
	return fmt.Errorf("bad verification code")
}

// SendCode ...
func (inst *VerificationServiceImpl) SendCode(c context.Context, v *services.Verification) error {
	if v.Code == "" {
		v.Code = inst.makeVerificationCode()
	}
	err := inst.dispatchMessage(c, v)
	if err != nil {
		return err
	}
	vlog.Debug("vcode: %s", v.Code)
	return inst.makeToken(c, v)
}

func (inst *VerificationServiceImpl) getAccount(v *services.Verification) string {
	mech := v.Mechanism
	switch mech {
	case auth.MechanismSMS:
		return v.ToPhone.String()
	case auth.MechanismEmail:
		return v.ToMail.String()
	default:
	}
	return ("bad mechanism:" + mech)
}

func (inst *VerificationServiceImpl) makeUUID() lang.UUID {
	b := inst.UUIDService.Build()
	b.Class("VerificationCode")
	return b.Generate()
}

func (inst *VerificationServiceImpl) makeSalt(t lang.Time) lang.Base64 {
	return inst.saltGen.salt(t)
}

func (inst *VerificationServiceImpl) makeComputeT2(t1 lang.Time) lang.Time {
	delta := lang.Time(inst.TokenMaxAgeInMS)
	return t1 + delta
}

func (inst *VerificationServiceImpl) makeToken(c context.Context, v *services.Verification) error {

	t1 := lang.Now()
	t2 := inst.makeComputeT2(t1)
	uuid := inst.makeUUID()
	salt := inst.makeSalt(t1)

	// get token
	token, err := inst.JWT.GetDTO(c)
	if err != nil {
		token = &jwt.Token{}
	}

	// data1
	data1 := &data4sum{}
	data1.code = v.Code
	data1.mechanism = v.Mechanism
	data1.account = inst.getAccount(v)
	data1.uuid = uuid
	data1.salt = salt
	data1.t1 = t1
	data1.t2 = t2

	// data2
	data2 := &data4jwt{
		t1:   t1,
		t2:   t2,
		uuid: uuid,
	}
	sum := data1.computeHash()
	data2.sum = lang.Base64FromBytes(sum)

	// set token
	token.Properties = data2.saveProperties(token.Properties)
	token.StartedAt = t1
	token.ExpiredAt = t2
	return inst.JWT.SetDTO(c, token)
}

func (inst *VerificationServiceImpl) dispatchMessage(c context.Context, v *services.Verification) error {

	text := "Your Verification Code is : " + v.Code
	var to mails.Address
	var from mails.Address
	mech := v.Mechanism

	if mech == "" {
		return fmt.Errorf("bad mechanism: " + mech)
	} else if mech == auth.MechanismSMS {
		from = mails.Address(inst.SenderFromSMS)
		to = mails.Address(v.ToPhone)
	} else if mech == auth.MechanismEmail {
		from = mails.Address(inst.SenderFromMail)
		to = mails.Address(v.ToMail)
	} else {
		return fmt.Errorf("bad mechanism: " + mech)
	}

	msg := &mails.Message{}
	msg.Title = "Your Verification Code"
	msg.FromAddress = from
	msg.ToAddresses = []mails.Address{to}
	msg.ContentType = "text/plain"
	msg.Content = []byte(text)

	return inst.Sender.Send(c, msg)
}

func (inst *VerificationServiceImpl) makeVerificationCode() string {
	const size = 7
	uuid := inst.UUIDService.Build().Class("v-code").Generate()
	str := uuid.String()
	src := []rune(str)
	b := strings.Builder{}
	for _, ch := range src {
		if b.Len() < size {
			if '0' <= ch && ch <= '9' {
				b.WriteRune(ch)
			}
		} else {
			return b.String()
		}
	}
	for b.Len() < size {
		b.WriteRune('0')
	}
	return b.String()
}
