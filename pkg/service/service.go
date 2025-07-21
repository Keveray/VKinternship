package service

import repository "github.com/Keveray/VKinternship/pkg/repository"

type Authorization interface {
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
	return &Service{}
}
