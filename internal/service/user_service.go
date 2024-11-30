package service

import (
	"time"

	"github.com/aidosgal/gust/internal/dto"
	hashlib "github.com/aidosgal/gust/internal/lib/hash"
	jwtlib "github.com/aidosgal/gust/internal/lib/jwt"
	"github.com/aidosgal/gust/internal/model"
)

type UserRepository interface {
	GetUserByPhone(phone string, app_id int) (model.User, error)
	GetAppByUser(app_id int) (model.App, error)
	GetUser(user_id int) (model.User, error)
	CreateUser(req dto.RegisterRequest) (int, error)
}

type UserService struct {
	repository UserRepository
}

func NewUserService(repository UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) Login(req dto.LoginRequest) (string, error) {
	user, err := s.repository.GetUserByPhone(req.Phone, req.AppId)
	if err != nil {
		return "", err
	}

	if !hashlib.CheckPasswordHash(user.Password, req.Password) {
		return "", err
	}

	app, err := s.repository.GetAppByUser(req.AppId)
	if err != nil {
		return "", err
	}

	token, err := jwtlib.NewToken(user, app, time.Hour*24*365)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *UserService) Register(req dto.RegisterRequest) (string, error) {
	hashPassword, err := hashlib.HashUserPassword(req.Password)
	if err != nil {
		return "", err
	}

	req.Password = hashPassword

	user_id, err := s.repository.CreateUser(req)
	if err != nil {
		return "", err
	}

	app, err := s.repository.GetAppByUser(req.AppId)
	if err != nil {
		return "", err
	}

	user, err := s.repository.GetUser(user_id)
	if err != nil {
		return "", err
	}

	token, err := jwtlib.NewToken(user, app, time.Hour*24*365)
	if err != nil {
		return "", err
	}
	return token, nil
}
