package usecase

import (
	userDomain "api/domain/user"
	"context"
)

type FindUserUseCase struct {
	ur userDomain.UserRepository
}
type FindUserUseCaseOutputDTO struct {
	ID    string
	Email string
	Type  string
}

func NewFindUserUseCase(ur userDomain.UserRepository) *FindUserUseCase {
	return &FindUserUseCase{ur}
}

func (uc *FindUserUseCase) Run(ctx context.Context) (*FindUserUseCaseOutputDTO, error) {
	user, err := uc.ur.Find()
	if err != nil {
		return nil, err
	}
	return &FindUserUseCaseOutputDTO{
		ID:    user.ID,
		Email: user.Email,
		Type:  user.Type,
	}, nil
}
