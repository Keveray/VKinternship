package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/Keveray/VKinternship"
	"github.com/Keveray/VKinternship/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	salt = "u2f3fi2hnh2jhf2jf23jhf"
	tokenTTL = 12 * time.Hour
	signinKey = "nhiocI234YC7832NICH4I2HL"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

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

func (s AuthService) GenerateToken(username, password string) (int, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password)))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		}
		user.Id,
	})
	return token.SignedString([]byte(signinKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))

}
