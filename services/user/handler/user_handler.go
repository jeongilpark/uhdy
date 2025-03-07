package handler

import (
	"context"

	"github.com/danielgtaylor/huma/v2"

	"sagepulse.ai/uhdy/user-service/model"
	"sagepulse.ai/uhdy/user-service/repository"
	"sagepulse.ai/uhdy/user-service/service"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(repo repository.UserRepository, jwtKey string) (*UserHandler, error) {
	userService, err := service.NewUserService(repo, jwtKey)
	if err != nil {
		return nil, err
	}
	return &UserHandler{userService: userService}, nil
}

// signUp handles user registration
func (h *UserHandler) SignUp(ctx context.Context, input *model.AuthRequest) (*struct{}, error) {
	err := h.userService.SignUp(ctx, input.Body)
	if err == nil {
		return nil, nil
	}
	return nil, huma.Error500InternalServerError(err.Error())
}

// signIn handles user login
func (h *UserHandler) SignIn(ctx context.Context, input *model.AuthRequest) (*model.SignInResponse, error) {
	token, err := h.userService.SignIn(ctx, input.Body)
	if err == nil {
		body := model.AuthToken{Token: token}
		return &model.SignInResponse{Body: body}, nil
	}
	if err == repository.ErrNoRecord {
		return nil, huma.Error404NotFound("Not found")
	}
	return nil, huma.Error500InternalServerError(err.Error())
}
