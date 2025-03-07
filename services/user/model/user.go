package model

import (
	"github.com/dgrijalva/jwt-go"
)

//
// Models for API
//

// Credentials for SignUp/SignIn
type Credentials struct {
	Username string `json:"username" required:"true" example:"account@example.com" doc:"Unique user name, such like email for example"`
	Password string `json:"password" required:"true" doc:"Password"`
}

type AuthToken struct {
	Token string `json:"token" required:"true" doc:"JWT token"`
}

//
// Models for API Interface
//

type AuthRequest struct {
	Body Credentials
}

type SignInResponse struct {
	Body AuthToken
}

//
// Models for the internal usage
//

// Claims in JWT Auth Token
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
