package repository

import (
	"github.com/Keveray/VKinternship"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user VKinternship.User) (int, error)
}

type VkInternshipList interface {
}

type VkInternshipItem interface {
}

type Repository struct {
	Authorization
	VkInternshipList
	VkInternshipItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
