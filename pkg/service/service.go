package service

import "github.com/kagorbunov/user_balance/pkg/repository"

type Users interface {
}

type Service struct {
	Users
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
