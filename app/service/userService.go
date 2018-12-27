package service

import (
	"context"
	"github.com/rojoherrero/quality-accounts/app/model"
	"github.com/rojoherrero/quality-accounts/app/repository"
	"github.com/rojoherrero/quality-common"
	"golang.org/x/crypto/bcrypt"
)

type (
	AccountService interface {
		Save(ctx context.Context, user model.UserCreationDto) error
		Update(ctx context.Context, user model.UserCreationDto) error
		Paginate(ctx context.Context, start, end int64) (model.PropertyMapSlice, error)
		Delete(ctx context.Context, id int64) error
		GetLoginData(username, password string) model.User
	}

	accountService struct {
		repo   repository.UserRepository
	}
)

func NewAccountService(repo repository.UserRepository) AccountService {
	return &accountService{repo}
}

func (s *accountService) Save(ctx context.Context, user model.UserCreationDto) error {
	hashedPwd, e := s.hashPassword(user.Password)
	if e != nil {
		return e
	}
	user.Password = hashedPwd
	return s.repo.Save(ctx, user)
}

func (s *accountService) Update(ctx context.Context, user model.UserCreationDto) error {
	hashedPwd, e := s.hashPassword(user.Password)
	if e != nil {
		return e
	}
	user.Password = hashedPwd
	return s.repo.Update(ctx, user)
}

func (s *accountService) hashPassword(password string) (string, error) {
	hash, e := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), e
}

func (s *accountService) Paginate(ctx context.Context,start, end int64) (model.PropertyMapSlice, error) {
	return s.repo.Paginate(ctx, start, end)
}

func (s *accountService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

func (s *accountService) GetLoginData(username, password string) model.User {
	return nil
}
