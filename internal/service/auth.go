package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"testbackendGraudate/domain"
	"testbackendGraudate/internal/repository"
	"time"
)

const (
	passwordSalt = "sdf209kjswfy824fhd>?@1$12063512a>d;239fdsfad41"
	tokenTTL     = 12 * time.Hour
	signingKey   = "fkdshf743fhufs1qfr61.d32r2!!23,<#>#!1247"
)

type tokenClaims struct {
	jwt.MapClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user domain.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(email, password string) (string, error) {
	user, err := s.repo.LoginUser(email, s.generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		jwt.MapClaims{
			"exp": time.Now().Add(tokenTTL).Unix(),
			"iat": time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(passwordSalt)))
}
