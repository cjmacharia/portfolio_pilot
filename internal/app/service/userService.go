package service

import (
	"github.com/cjmacharia/portfolioPilot/internal/domain"
)

type UserService interface {
	UserSignUp(u *domain.User) (int, error)
	GetUserByIDService(email string) (domain.User, error)
	GetUserWalletBalanceService(userID int) (float64, error)
	UpdateUserWalletBalanceService(userID int, newBalance float64) error
}

type DefaultUserService struct {
	repo domain.UserRepository
}

func (us DefaultUserService) UserSignUp(u *domain.User) (int, error) {
	return us.repo.SignUp(u)
}

func (us DefaultUserService) GetUserByIDService(email string) (domain.User, error) {
	return us.repo.GetUserByEmail(email)
}

func (us DefaultUserService) GetUserWalletBalanceService(userID int) (float64, error) {
	return us.repo.GetUserWalletBalance(userID)
}
func (us DefaultUserService) UpdateUserWalletBalanceService(userID int, newBalance float64) error {
	return us.repo.UpdateUserWalletBalance(userID, newBalance)
}

func NewUserService(repository domain.UserRepository) DefaultUserService {
	return DefaultUserService{
		repository,
	}
}
