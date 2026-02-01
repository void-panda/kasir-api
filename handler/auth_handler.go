package handler

import (
	"kasir-api/model"
	"kasir-api/service"
	"kasir-api/utils"
	"net/http"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(srvc *service.AuthService) *AuthHandler {
	return &AuthHandler{service: srvc}
}

// Register godoc
// @Summary Register new user
// @Description Register new user with Name, Email and Password
// @Tags auth
// @Accpet json
// @Produce json
// @Param user body model.RegisterRequest true "New User Credentials" SchemaExample({"name":"User Name", "email": "user@email.example", "password": "#password123"})
// @Success 201 {object} model.Response
// @Router /api/auth/register [post]
func (hdlr *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req model.RegisterRequest

	err := utils.BindAndValidate(r, &req)
	if err != nil {
		model.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := hdlr.service.Register(req.Name, req.Email, req.Password)
	if err != nil {
		model.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	user.Password = ""
	model.Success(w, http.StatusCreated, "New user registered succesfully. Please login", user)
}

// Login godoc
// @Summary Login
// @Description Login by using Email and Password
// @Tags auth
// @Accpet json
// @Produce json
// @Param user body model.LoginRequest true "New User Credentials" SchemaExample("email": "user@email.example", "password": "#password123"})
// @Success 200 {object} model.Response
// @Router /api/auth/login [post]
func (hdlr *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req model.LoginRequest

	err := utils.BindAndValidate(r, &req)
	if err != nil {
		model.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	user, token, err := hdlr.service.Login(req.Email, req.Password)
	if err != nil {
		model.Error(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	user.Password = ""

	model.Success(w, http.StatusOK, "Login success", model.LoginResponse{User: user, Token: token})

}
