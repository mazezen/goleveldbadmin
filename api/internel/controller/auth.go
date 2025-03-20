package controller

import (
	"encoding/json"
	"github.com/mazezen/goleveldbadmin/framework"
	"github.com/mazezen/goleveldbadmin/service"
	"net/http"
)

var authService = new(service.AuthService)

/**
 * r.PostFormValue  : 可以解析 Post/PUT Content-Type=application/x-www-form-urlencoded 或 Content-Type=multipart/form-data
 */

type AuthHandler struct{}

func (p *AuthHandler) Router(router *framework.RouterHandler) {
	router.Router("/lb/login", p.login)
}

type loginPayload struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func (p *AuthHandler) login(w http.ResponseWriter, r *http.Request) {
	cors(w, r)
	result := &framework.Result{
		Code:    2000,
		Message: "登录成功",
	}
	var payload loginPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		result.Code = 5000
		result.Message = err.Error()
		framework.JsonOk(w, result)
		return
	}

	token, err := authService.Login(payload.Account, payload.Password)
	if err != nil {
		result.Code = 5000
		result.Message = err.Error()
		return
	}
	result.Data = token

	framework.JsonOk(w, result)
}
