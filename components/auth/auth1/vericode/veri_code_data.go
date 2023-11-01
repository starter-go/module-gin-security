package vericode

import (
	"bytes"
	"crypto/sha256"
	"strconv"

	"github.com/starter-go/base/lang"
)

// 用于读写 jwt
type data4jwt struct {
	sum  lang.Base64 // hash 值
	t1   lang.Time   // 有效期（开始）
	t2   lang.Time   // 有效期（截止）
	uuid lang.UUID   // 标识符
}

// 用于计算 hash
type data4sum struct {
	// sum  lang.Base64

	t1   lang.Time
	t2   lang.Time
	uuid lang.UUID

	code      string
	account   string
	mechanism string
	salt      lang.Base64 // 这个盐参数由 rand_loop(t1) 计算生成
}

////////////////////////////////////////////////////////////////////////////////

type timeString struct{}

func (x *timeString) stringify(t lang.Time) string {
	n := t.Int()
	return strconv.FormatInt(n, 10)
}

func (x *timeString) parse(str string) lang.Time {
	n, _ := strconv.ParseInt(str, 10, 64)
	return lang.Time(n)
}

////////////////////////////////////////////////////////////////////////////////

const (
	data4jwtKeyPrefix = "vericode."

	data4jwtKeyT1   = data4jwtKeyPrefix + "t1"
	data4jwtKeyT2   = data4jwtKeyPrefix + "t2"
	data4jwtKeySUM  = data4jwtKeyPrefix + "sum"
	data4jwtKeyUUID = data4jwtKeyPrefix + "uuid"
)

func (inst *data4jwt) loadProperties(props map[string]string) {
	if props == nil {
		return
	}

	t1 := props[data4jwtKeyT1]
	t2 := props[data4jwtKeyT2]
	sum := props[data4jwtKeySUM]
	uuid := props[data4jwtKeyUUID]

	ts := timeString{}
	inst.t1 = ts.parse(t1)
	inst.t2 = ts.parse(t2)
	inst.sum = lang.Base64(sum)
	inst.uuid = lang.ParseUUID(uuid)
}

func (inst *data4jwt) saveProperties(props map[string]string) map[string]string {
	if props == nil {
		props = make(map[string]string)
	}

	ts := timeString{}
	t1 := ts.stringify(inst.t1)
	t2 := ts.stringify(inst.t2)
	sum := inst.sum.String()
	uuid := inst.uuid.String()

	props[data4jwtKeyT1] = t1
	props[data4jwtKeyT2] = t2
	props[data4jwtKeySUM] = sum
	props[data4jwtKeyUUID] = uuid

	return props
}

////////////////////////////////////////////////////////////////////////////////

func (inst *data4sum) computeHash() []byte {

	ts := timeString{}

	list := make([]string, 0)
	list = append(list, inst.account)
	list = append(list, inst.code)
	list = append(list, inst.mechanism)
	list = append(list, ts.stringify(inst.t1))
	list = append(list, ts.stringify(inst.t2))
	list = append(list, inst.uuid.String())

	b := bytes.Buffer{}
	for _, item := range list {
		b.WriteString(item)
		b.WriteByte(0)
	}
	salt := inst.salt.Bytes()
	b.Write(salt)

	sum := sha256.Sum256(b.Bytes())
	return sum[:]
}

////////////////////////////////////////////////////////////////////////////////
