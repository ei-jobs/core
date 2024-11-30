package handler

import (
	"fmt"
	"net/http"

	"github.com/aidosgal/gust/internal/dto"
	jsonlib "github.com/aidosgal/gust/internal/lib/json"
	jwtlib "github.com/aidosgal/gust/internal/lib/jwt"
)

type UserHandler struct {
	service UserService
}

type UserService interface {
	Login(req dto.LoginRequest) (string, error)
	Register(req dto.RegisterRequest) (string, error)
	UpdateUser(req dto.UpdateRequest, token string) error
	DeleteUser(token string, app_id int) error
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest
	err := jsonlib.ParseJSON(r, &req)
	if err != nil {
		jsonlib.WriteError(w, http.StatusBadRequest, err)
		return
	}

	token, err := h.service.Login(req)
	if err != nil {
		jsonlib.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid credentials"))
		return
	}

	jsonlib.WriteJSON(w, http.StatusAccepted, map[string]interface{}{"token": token})
	return
}

func (h *UserHandler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterRequest
	err := jsonlib.ParseJSON(r, &req)
	if err != nil {
		jsonlib.WriteError(w, http.StatusBadRequest, err)
		return
	}

	token, err := h.service.Register(req)
	if err != nil {
		jsonlib.WriteError(w, http.StatusBadRequest, err)
		return
	}

	jsonlib.WriteJSON(w, http.StatusCreated, map[string]interface{}{"token": token})
	return
}

func (h *UserHandler) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	return
}

func (h *UserHandler) HandeGetMe(w http.ResponseWriter, r *http.Request) {
	return
}

func (h *UserHandler) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	token, err := jwtlib.GetToken(r)
	if err != nil {
		jsonlib.WriteError(w, http.StatusUnauthorized, err)
		return
	}

	var req dto.UpdateRequest
	err = jsonlib.ParseJSON(r, &req)
	if err != nil {
		jsonlib.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = h.service.UpdateUser(req, token)
	if err != nil {
		jsonlib.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	return
}

func (h *UserHandler) HandleDelete(w http.ResponseWriter, r *http.Request) {
	token, err := jwtlib.GetToken(r)
	if err != nil {
		jsonlib.WriteError(w, http.StatusUnauthorized, err)
		return
	}

	var req dto.DeleteRequest
	err = jsonlib.ParseJSON(r, &req)
	if err != nil {
		jsonlib.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = h.service.DeleteUser(token, req.AppId)
	if err != nil {
		jsonlib.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	return
}
