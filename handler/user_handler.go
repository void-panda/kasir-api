package handler

import (
	"encoding/json"
	"kasir-api/model"
	"kasir-api/service"
	"net/http"
	"strconv"
)

type UserHandler struct {
	service *service.UserService
}

// struct method for : "UserHandler" struct
func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// === handler Functions ===

// GetAll godoc
// @Summary Get All Users
// @Description Fetch all users data
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} model.Response
// @Router /api/users [get]
func (hdlr *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := hdlr.service.GetAll()
	if err != nil {
		model.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	model.Success(w, http.StatusOK, "successfully get users", users)
}

// GetById godoc
// @Summary Get User by ID
// @Description Get a single user data by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} model.Response
// @Failure 404 {object} model.Response
// @Router /api/users/{id} [get]
func (hdlr *UserHandler) GetById(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		model.Error(w, http.StatusBadRequest, "Invalid User ID")
	}

	user, err := hdlr.service.GetByID(id)
	if err != nil {
		model.Error(w, http.StatusNotFound, err.Error())
		return
	}
	model.Success(w, http.StatusOK, "succesfully get user", user)
}

// Update godoc
// @Summary Update user
// @Description Update produk berdasarkan ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body model.User true "User Data"
// @Success 200 {object} model.Response
// @Failure 404 {object} model.Response
// @Router /api/users/{id} [put]
func (hdlr *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		model.Error(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var user model.User
	// if err = utils.BindAndValidate(r, &user); err != nil {
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		model.Error(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	user.ID = id
	err = hdlr.service.Update(&user)
	if err != nil {
		model.Error(w, http.StatusNotFound, err.Error())
		return
	}

	model.Success(w, http.StatusOK, "successfully updated user", user)
}

// Delete godoc
// @Summary Delete user
// @Description Menghapus produk berdasarkan ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} model.Response
// @Failure 404 {object} model.Response
// @Router /api/users/{id} [delete]
func (hdlr *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		model.Error(w, http.StatusBadRequest, "Invalid User ID")
		return
	}

	err = hdlr.service.Delete(id)
	if err != nil {
		model.Error(w, http.StatusNotFound, err.Error())
		return
	}

	model.Success(w, http.StatusOK, "successfully delete user", nil)
}
