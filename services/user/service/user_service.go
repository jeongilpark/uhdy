package service

import (
	"context"
	"encoding/base64"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"sagepulse.ai/uhdy/user-service/model"
	"sagepulse.ai/uhdy/user-service/repository"
)

type UserService interface {
	SignUp(ctx context.Context, creds model.Credentials) error
	SignIn(ctx context.Context, creds model.Credentials) (string, error)
}

type userService struct {
	repo   repository.UserRepository
	jwtKey []byte
}

func NewUserService(repo repository.UserRepository, jwtKey string) (UserService, error) {
	decoded_key, err := base64.StdEncoding.DecodeString(jwtKey)
	if err != nil {
		return nil, err
	}
	return &userService{repo: repo, jwtKey: decoded_key}, nil
}

func (s *userService) SignUp(ctx context.Context, creds model.Credentials) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return s.repo.CreateUser(ctx, strings.ToLower(creds.Username), string(hashedPassword))
}

func (s *userService) SignIn(ctx context.Context, creds model.Credentials) (string, error) {
	user, err := s.repo.GetUser(ctx, strings.ToLower(creds.Username))
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(creds.Password))
	if err != nil {
		return "", err
	}

	expirationTime := time.Now().UTC().Add(5 * time.Minute)
	claims := &model.Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			Subject:   user.ID.String(),
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
