package service

import (
	"cryptoserver/internal/repository"
	"errors"
)

type AuthService struct {
	userRepository *repository.UserRep
}

func NewAuthService(userRepository *repository.UserRep) *AuthService {
	return &AuthService{
		userRepository: userRepository,
	}
}

func (s *AuthService) RegisterUser(name string, password string) error {

	//1.Проверить есть ли пользователь. Если есть то отменить регистрацию
	//2.Создать и Добавить нового пользователя в мапу
	check := s.userRepository.ExistByUserName(name)

	if check {
		return errors.New("user already exist")
	}

	user := repository.User{
		Username: name,
		Password: password,
	}

	//добавление и создание в мапу
	if err := s.userRepository.Create(&user); err != nil {
		return err
	}

	return nil

}
