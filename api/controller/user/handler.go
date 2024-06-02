package user

import (
	"api/usecase"
	"context"
	"net/http"
)

type Handler struct {
	FindUserUseCase *usecase.FindUserUseCase
}

func NewHandler(fu *usecase.FindUserUseCase) Handler {
	return Handler{
		FindUserUseCase: fu,
	}
}

func (h Handler) Get(ctx context.Context, r *http.Request) (*userResponseModel, error) {
	dtos, err := h.FindUserUseCase.Run(ctx)
	if err != nil {
		return nil, err
	}
	return &userResponseModel{
		ID:    dtos.ID,
		Email: dtos.Email,
	}, nil
}
