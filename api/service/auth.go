package service

import (
	"errors"
	"github.com/mazezen/goleveldbadmin/sdk"
	"github.com/mazezen/goleveldbadmin/utils"
)

type AuthService struct{}

var NewAuthService = new(AuthService)

func (a *AuthService) Login(account, password string) (string, error) {
	if account == "" || password == "" {
		return "", errors.New("account or password is empty")
	}

	if account != sdk.Cf.Ti.Acc || utils.Sha256(password) != sdk.Cf.Ti.Pwd {
		return "", errors.New("账号和密码错误")
	}
	// sign token
	claims := &utils.JWTClaims{
		Account: account,
	}
	token, err := utils.GenToken(claims)
	if err != nil {
		return "", err
	}
	return token, nil
}
