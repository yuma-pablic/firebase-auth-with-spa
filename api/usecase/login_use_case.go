package usecase

import (
	userDomain "api/domain/user"
	"context"
)

type LoginUserUseCase struct {
	ur userDomain.UserRepository
}
type LoginUserUseCaseOutputDTO struct {
	ID    string
	Email string
	Type  string
}

func NewLoginUserUseCase(ur userDomain.UserRepository) *LoginUserUseCase {
	return &LoginUserUseCase{ur}
}

func (uc *LoginUserUseCase) Run(ctx context.Context) (*FindUserUseCaseOutputDTO, error) {
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
