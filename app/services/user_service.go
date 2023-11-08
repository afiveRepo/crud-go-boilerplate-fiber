package services

import (
	"crud-go-boilerplate-fiber/app/helper"
	"crud-go-boilerplate-fiber/app/libraries/mail"
	"crud-go-boilerplate-fiber/app/libraries/middleware"
	"crud-go-boilerplate-fiber/app/models/entities"
	"crud-go-boilerplate-fiber/app/models/requests"
	"crud-go-boilerplate-fiber/app/models/responses"
	"crud-go-boilerplate-fiber/app/repository"
	"errors"

	"github.com/google/uuid"
)

type UserService interface {
	Create(req requests.UserRequest) (*responses.UsersResponse, error)
	Login(req requests.UserLoginRequest) (*responses.UsersResponse, error)
}
type userService struct {
	base repository.BaseRepository
	user repository.UserRepository
}

func NewUserService(
	base repository.BaseRepository,
	user repository.UserRepository,
) UserService {
	return &userService{
		base: base,
		user: user,
	}
}
func (s *userService) Create(req requests.UserRequest) (*responses.UsersResponse, error) {
	uid := uuid.New().String()
	password, _ := helper.GetAESEncrypted(req.Password)
	newUser := entities.User{
		Uid:         uid,
		Email:       req.Email,
		Password:    password,
		Fullname:    req.Fullname,
		PhoneNumber: req.PhoneNumber,
	}
	tx := s.base.GetBegin()
	user, err := s.user.Create(tx, newUser)
	if err != nil {
		return nil, err
	}
	token, err := middleware.GenerateToken(user.ID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	to := []string{
		user.Email,
	}
	errSend := mail.SendEmail(to, nil, "testing", "hallo gasa")
	if errSend != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	result := responses.UsersResponse{
		Fullname: user.Fullname,
		Email:    user.Email,
		Token:    token,
	}
	return &result, nil
}
func (s *userService) Login(req requests.UserLoginRequest) (*responses.UsersResponse, error) {

	password, _ := helper.GetAESEncrypted(req.Password)
	user, err := s.user.FindbyEmail(req.Email)
	if err != nil {
		return nil, err
	}
	if user.Password != password {
		errs := errors.New("wrong password")
		return nil, errs
	}
	token, err := middleware.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}
	result := responses.UsersResponse{
		Fullname: user.Fullname,
		Email:    user.Email,
		Token:    token,
	}
	return &result, nil
}
