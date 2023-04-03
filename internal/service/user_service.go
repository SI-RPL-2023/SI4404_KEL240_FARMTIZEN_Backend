package service

import (
	"github.com/SI-RPL-2023/SI4404_KEL240_FARMTIZEN_Backend/internal/domain/entity"
	"github.com/SI-RPL-2023/SI4404_KEL240_FARMTIZEN_Backend/internal/domain/repository"
)

type UserService struct {
	userRepo repository.UserRepo
}

func NewUserService(userRepo repository.UserRepo) *UserService {
	return &UserService{userRepo: userRepo}
}

// SignUp is a function to register a new user
func (s *UserService) SignUp(user *entity.User) error {
	return nil
}
