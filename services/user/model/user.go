package model

import (
	"github.com/dgrijalva/jwt-go"
)

// Credentials for SignUp/SignIn
type Credentials struct {
	Username string `json:"username" required:"true" example:"account@example.com" doc:"Unique user name, such like email for example"`
	Password string `json:"password" required:"true" doc:"Password"`
}

type AuthInput struct {
	Body Credentials
}

type AuthToken struct {
	Token string `json:"token" required:"true" doc:"JWT token"`
}

type SignInResponse struct {
	Body AuthToken
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
