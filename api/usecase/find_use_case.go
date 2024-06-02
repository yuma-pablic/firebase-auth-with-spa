package usecase

import (
	userDomain "api/domain/user"
	"context"
)

type FindUserUseCase struct {
	ur userDomain.UserRepository
}

func NewFindUserUseCase(ur userDomain.UserRepository) *FindUserUseCase {
	return &FindUserUseCase{ur}
}

func (uc *FindUserUseCase) Run(ctx context.Context) (*userDomain.User, error) {
	user, err := uc.ur.Find()
	if err != nil {
		return nil, err
	}
	return user, nil
}
