package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/Keveray/VKinternship"
	"github.com/Keveray/VKinternship/pkg/repository"
)

const salt = "u2f3fi2hnh2jhf2jf23jhf"

type AuthService struct {
	repo repository.Authorization
}

func NewAutoService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user VKinternship.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))

}
