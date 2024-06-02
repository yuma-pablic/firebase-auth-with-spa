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

func (uc *LoginUserUseCase) Run(ctx context.Context) (*LoginUserUseCaseOutputDTO, error) {
	user, err := uc.ur.Find()
	if err != nil {
		return nil, err
	}
	return &LoginUserUseCaseOutputDTO{
		ID:    user.ID,
		Email: user.Email,
		Type:  user.Type,
	}, nil
}
