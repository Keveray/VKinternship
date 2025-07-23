package service

import (
	"github.com/Keveray/VKinternship"
	repository "github.com/Keveray/VKinternship/pkg/repository"
)

type Authorization interface {
	CreateUser(user VKinternship.User) (int, error)
	GenerateToken(username, password string) (int, error)
}

type VkInternshipList interface {
}

type VkInternshipItem interface {
}

type Service struct {
	Authorization
	VkInternshipList
	VkInternshipItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
