package service

import (
	"main/internal/app/model"
	"main/internal/app/store"
)

type UserService struct {
	userRepository *store.UserRepository
}

func CreateUserService(s *store.Store) *UserService {
	return &UserService{
		userRepository: &store.UserRepository{
			Store: s,
		},
	}
}

func (us *UserService) CreateUser(u *model.User) (*model.User, error) {
	user, err := us.userRepository.Create(u)

	if err != nil {
		return &model.User{}, err
	}

	return user, nil
}

func (us *UserService) FindUserByEmail(email string) (*model.User, error) {
	byEmail, err := us.userRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	return byEmail, nil
}
