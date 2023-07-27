package tokens

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/security/jwt"
)

const (
	headerJWT    = "x-jwt"
	headerSetJWT = "x-set-jwt"
)

// GinJWTokenAdapter ...
type GinJWTokenAdapter struct {

	//starter:component
	_as func(jwt.Adapter, jwt.Registry) //starter:as(".",".")

	JWTser jwt.Service //starter:inject("#")

	CookieMaxAge int  //starter:inject("${jwt.gin-adapter.cookie-max-age}")
	UseCookie    bool //starter:inject("${jwt.gin-adapter.use-cookie}")
	UseHeader    bool //starter:inject("${jwt.gin-adapter.use-header}")

}

func (inst *GinJWTokenAdapter) _impl() (jwt.Adapter, jwt.Registry) {
	return inst, inst
}

// ListRegistrations ...
func (inst *GinJWTokenAdapter) ListRegistrations() []*jwt.Registration {
	r1 := &jwt.Registration{
		Enabled:  true,
		Priority: 0,
		Adapter:  inst,
	}
	return []*jwt.Registration{r1}
}

func (inst *GinJWTokenAdapter) getGinContext(c context.Context) (*gin.Context, error) {
	gc, ok := c.(*gin.Context)
	if ok && (gc != nil) {
		return gc, nil
	}
	return nil, fmt.Errorf("the context is not a gin.Context")
}

// Accept ...
func (inst *GinJWTokenAdapter) Accept(c context.Context) bool {
	gc, _ := inst.getGinContext(c)
	return (gc != nil)
}

// GetDTO ...
func (inst *GinJWTokenAdapter) GetDTO(c context.Context) (*jwt.DTO, error) {
	text, err := inst.GetText(c)
	if err != nil {
		return nil, err
	}
	return inst.JWTser.Decode(text)
}

// GetText ...
func (inst *GinJWTokenAdapter) GetText(c context.Context) (jwt.Text, error) {

	gc, err := inst.getGinContext(c)
	if err != nil {
		return "", err
	}

	const minTextLen = 10
	name := headerJWT

	if inst.UseHeader {
		value := gc.GetHeader(name)
		if len(value) > minTextLen {
			return jwt.Text(value), nil
		}
	}

	if inst.UseCookie {
		value, err := gc.Cookie(name)
		if err == nil && len(value) > minTextLen {
			return jwt.Text(value), nil
		}
	}

	return "", fmt.Errorf("no JWT info in HTTP headers")
}

// SetDTO ...
func (inst *GinJWTokenAdapter) SetDTO(c context.Context, o *jwt.DTO) error {
	text, err := inst.JWTser.Encode(o)
	if err != nil {
		return err
	}
	return inst.SetText(c, text)
}

// SetText ...
func (inst *GinJWTokenAdapter) SetText(c context.Context, t jwt.Text) error {

	gc, err := inst.getGinContext(c)
	if err != nil {
		return err
	}

	value := t.String()

	if inst.UseCookie {
		name := headerJWT
		maxAge := inst.CookieMaxAge // in seconds
		path := ""
		domain := ""
		secure := false
		httpOnly := false
		gc.SetCookie(name, value, maxAge, path, domain, secure, httpOnly)
	}

	if inst.UseHeader {
		name := headerSetJWT
		gc.Header(name, value)
	}

	return nil
}
