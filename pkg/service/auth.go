package service

import (
	"crypto/sha1"
	"errors"
	"fmt"

	"ozinshe/pkg/repository"
	"ozinshe/schemas"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	salt     = "PCl9kW51SL4VQ"
	signKey  = "PCl9kwWSPEC5512lke5cSLWsd54VQ"
	tokenTTL = time.Hour * 24 * 30
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId  int  `json:"user_id"`
	IsAdmin bool `json:"is_admin"`
}

type AuthService struct {
	repos repository.Authorizhation
}

func NewAuthService(repository repository.Authorizhation) *AuthService {
	return &AuthService{repos: repository}
}

func (s *AuthService) CreateUser(user schemas.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repos.CreateUser(user)
}

func (s *AuthService) GenerateToken(email, password string) (string, error) {
	user, err := s.repos.GetUser(email, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
		user.IsAdmin,
	})

	return token.SignedString([]byte(signKey))
}

func (s *AuthService) ParseToken(accessToken string) (map[string]any, error) {
	user := make(map[string]any, 2)
	token, err := jwt.ParseWithClaims(
		accessToken,
		&tokenClaims{},
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid signing method")
			}
			return []byte(signKey), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return nil, errors.New("token claims are not of type *tokenClaims")
	}

	user["userId"] = claims.UserId
	user["isAdmin"] = claims.IsAdmin

	return user, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
