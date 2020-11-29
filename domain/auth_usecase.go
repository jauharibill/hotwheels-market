package domain

import (
	"context"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	"lelang-hotwheels/model"
)

type AuthInterface interface {
	Login() (error)
	Register(user model.UserModel) error
	ForgotPassword() error
}

type AuthUsecase struct {
	repo rel.Repository
	ctx context.Context
}

var authusecase = AuthUsecase{}

func (_a *AuthUsecase) Login(authReq model.AuthModel) error {
	return _a.repo.Find(_a.ctx,
		&authReq, where.Eq("username", authReq.Username),
		where.Eq("password", authReq.Password))
}

func (_a *AuthUsecase) Register(userReq model.UserModel) error {
	return _a.repo.Transaction(_a.ctx, func(ctx context.Context) error {
		return _a.repo.Insert(_a.ctx, &userReq)
	})
}

func (_a *AuthUsecase) ForgotPassword() error {
	// logic forgor password
	return nil
}
