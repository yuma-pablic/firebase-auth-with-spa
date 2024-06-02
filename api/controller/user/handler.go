package user

import (
	"api/usecase"
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type UserHandler interface {
	Get(ctx context.Context, r *http.Request) error
	Login(ctx context.Context, r *http.Request) error
	Logout(ctx context.Context, r *http.Request) error
	Delete(ctx context.Context, r *http.Request) error
}
type userHandler struct {
	FindUserUseCase *usecase.FindUserUseCase
}

func NewHandler(fu *usecase.FindUserUseCase) userHandler {
	return userHandler{
		FindUserUseCase: fu,
	}
}

func (h userHandler) Get(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	dtos, err := h.FindUserUseCase.Run(ctx)
	if err != nil {
		return err
	}
	if err != nil {
		if errors.As(err, &err) {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return err
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dtos)
	return nil
}

func (h userHandler) Login(ctx context.Context, r *http.Request) error {
	return nil
}
