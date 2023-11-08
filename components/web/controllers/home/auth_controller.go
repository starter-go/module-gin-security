package home

import (
	"github.com/gin-gonic/gin"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/libgin"
	"github.com/starter-go/security/auth"
	"github.com/starter-go/security/rbac"
)

// AuthVO ...
type AuthVO struct {
	rbac.AuthVO

	// rbac.BaseVO
	// Auth        rbac.AuthDTO  `json:"auth"`         // 需要验证的内容

	UserInfo    *rbac.UserDTO `json:"user"`         // 用户信息
	NewPassword lang.Base64   `json:"new_password"` // 新的密码（用于注册，设置，重设密码）
	Success     bool          `json:"success"`      // 是否完全成功

}

////////////////////////////////////////////////////////////////////////////////

// AuthController ...
type AuthController struct {

	//starter:component()
	_as func(libgin.Controller) //starter:as(".")

	Responder libgin.Responder //starter:inject("#")
	Service   rbac.AuthService //starter:inject("#")

	// AuthService services.AuthService //--starter:inject("#")

	AuthSer rbac.AuthService //starter:inject("#")

}

func (inst *AuthController) _impl() {
	inst._as(inst)
}

// Registration ...
func (inst *AuthController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{Route: inst.route}
}

func (inst *AuthController) route(g libgin.RouterProxy) error {
	g = g.For("auth")

	g.POST("", inst.handleSignIn)                                     // 'sign-in' 的简要别名
	g.POST("sign-in", inst.handleSignIn)                              // 登录
	g.POST("sign-up", inst.handleSignUp)                              // 注册
	g.POST("set-password", inst.handleSetPassword)                    // 修改密码
	g.POST("reset-password", inst.handleResetPassword)                // 重设密码
	g.POST("send-verification-mail", inst.handleSendVerificationMail) // 发送验证码(通过邮件)
	g.POST("send-verification-sms", inst.handleSendVerificationSMS)   // 发送验证码(通过短信)

	return nil
}

func (inst *AuthController) execute(req *myAuthRequest, fn func() error) {
	err := req.open()
	if err == nil {
		err = fn()
	}
	req.send(err)
}

func (inst *AuthController) handleSignIn(c *gin.Context) {
	req := &myAuthRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: true,
	}
	inst.execute(req, req.doLogin)
}

func (inst *AuthController) handleSignUp(c *gin.Context) {
	req := &myAuthRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: true,
	}
	inst.execute(req, req.doSignUp)
}

func (inst *AuthController) handleSendVerificationMail(c *gin.Context) {
	req := &myAuthRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: true,
	}
	inst.execute(req, req.doSendMail)
}

func (inst *AuthController) handleSendVerificationSMS(c *gin.Context) {
	req := &myAuthRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: true,
	}
	inst.execute(req, req.doSendSMS)
}

func (inst *AuthController) handleSetPassword(c *gin.Context) {
	req := &myAuthRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: true,
	}
	inst.execute(req, req.doSetPassword)
}

func (inst *AuthController) handleResetPassword(c *gin.Context) {
	req := &myAuthRequest{
		controller:      inst,
		context:         c,
		wantRequestBody: true,
	}
	inst.execute(req, req.doResetPassword)
}

////////////////////////////////////////////////////////////////////////////////

type myAuthRequest struct {
	// contexts
	controller *AuthController
	context    *gin.Context

	// flags
	wantRequestBody bool
	// wantRequestID   bool
	// wantRequestPage bool
	// wantRequestRBAC bool

	// params
	pagination rbac.Pagination
	id         rbac.PermissionID
	roles      rbac.RoleNameList

	// body
	body1 AuthVO
	body2 AuthVO
}

func (inst *myAuthRequest) open() error {

	c := inst.context

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myAuthRequest) send(err error) {
	resp := &libgin.Response{}
	resp.Data = &inst.body2
	resp.Context = inst.context
	resp.Status = inst.body2.Status
	resp.Error = err
	inst.controller.Responder.Send(resp)
}

func (inst *myAuthRequest) doSignUp() error {
	ctx := inst.context
	alist := inst.body1.Auth
	action := auth.ActionSignUp
	err := inst.controller.AuthSer.Handle(ctx, action, alist)
	return err
}

func (inst *myAuthRequest) doLogin() error {
	ctx := inst.context
	action := auth.ActionLogin
	alist := inst.body1.Auth
	err := inst.controller.AuthSer.Handle(ctx, action, alist)
	return err
}

func (inst *myAuthRequest) doSendMail() error {
	ctx := inst.context
	action := auth.ActionSendCode
	alist := inst.body1.Auth
	err := inst.controller.AuthSer.Handle(ctx, action, alist)
	return err
}

func (inst *myAuthRequest) doSendSMS() error {
	ctx := inst.context
	action := auth.ActionSendCode
	alist := inst.body1.Auth
	err := inst.controller.AuthSer.Handle(ctx, action, alist)
	return err
}

func (inst *myAuthRequest) doSetPassword() error {
	ctx := inst.context
	action := auth.ActionChangePassword
	alist := inst.body1.Auth
	err := inst.controller.AuthSer.Handle(ctx, action, alist)
	return err
}

func (inst *myAuthRequest) doResetPassword() error {
	ctx := inst.context
	action := auth.ActionResetPassword
	alist := inst.body1.Auth
	err := inst.controller.AuthSer.Handle(ctx, action, alist)
	return err
}

////////////////////////////////////////////////////////////////////////////////
