package service

//go:generate mockgen -source=$GOFILE -destination=../mock/mock_$GOFILE -package=mock

import (
	"context"
	"github.com/rojoherrero/quality-accounts/backend/model"
	"github.com/rojoherrero/quality-accounts/backend/repository"
	"golang.org/x/crypto/bcrypt"
)

type (
	UserService interface {
		Save(ctx context.Context, user model.UserCreationDto) error
		Update(ctx context.Context, user model.UserCreationDto) error
		Paginate(ctx context.Context, start, end int64) (model.PropertyMapSlice, error)
		Delete(ctx context.Context, id int64) error
	}

	userService struct {
		repo   repository.UserRepository
	}
)

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) Save(ctx context.Context, user model.UserCreationDto) error {
	hashedPwd, e := s.hashPassword(user.Password)
	if e != nil {
		return e
	}
	user.Password = hashedPwd
	return s.repo.Save(ctx, user)
}

func (s *userService) Update(ctx context.Context, user model.UserCreationDto) error {
	hashedPwd, e := s.hashPassword(user.Password)
	if e != nil {
		return e
	}
	user.Password = hashedPwd
	return s.repo.Update(ctx, user)
}

func (s *userService) hashPassword(password string) (string, error) {
	hash, e := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), e
}

func (s *userService) Paginate(ctx context.Context,start, end int64) (model.PropertyMapSlice, error) {
	return s.repo.Paginate(ctx, start, end)
}

func (s *userService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
