package handler

import (
	"fmt"
	"net/http"

	"github.com/aidosgal/gust/internal/dto"
	jsonlib "github.com/aidosgal/gust/internal/lib/json"
)

type UserHandler struct {
	service UserService
}

type UserService interface {
	Login(req dto.LoginRequest) (string, error)
	Register(req dto.RegisterRequest) (string, error)
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
