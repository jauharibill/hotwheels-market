package handler

import (
	"encoding/json"
	"lelang-hotwheels/domain"
	"lelang-hotwheels/model"
	"lelang-hotwheels/presenter"
	"net/http"
)

type AuthInterface interface {
	Home(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
	ForgotPassword(w http.ResponseWriter, r *http.Request)
}

type AuthHandler struct {
	usecase domain.AuthUsecase
}

var Auth = AuthHandler{}

func (_a *AuthHandler) Home(w http.ResponseWriter, r *http.Request) {
	var response presenter.Response

	response.Message = "this is home"
	output, _ := json.Marshal(response)

	w.Write(output)
}

func (_a *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var auth model.AuthModel
	var response presenter.Response

	json.NewDecoder(r.Body).Decode(&auth)

	err := _a.usecase.Login(auth)

	if err != nil {
		response.Message = err.Error()
	}

	response.Message = "Success Login"
	output, _ := json.Marshal(response)

	w.Write(output)
}

func (_a *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var response presenter.Response
	var user model.UserModel

	json.NewDecoder(r.Body).Decode(&user)

	err := _a.usecase.Register(user)

	if err != nil {
		response.Message = err.Error()
	}

	response.Message = "Success Register"
	output, _ := json.Marshal(response)

	w.Write(output)
}

func (_a *AuthHandler) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	
}
