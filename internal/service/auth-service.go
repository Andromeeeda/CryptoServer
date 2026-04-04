package service

import (
	"cryptoserver/internal/repository"
	"cryptoserver/pkg/jwt"
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

func (s *AuthService) RegisterUser(name string, password string) (string, error) {

	//1.Проверить есть ли пользователь. Если есть то отменить регистрацию
	//2.Создать и Добавить нового пользователя в мапу
	//Создать JWt токен и вернуть его
	check := s.userRepository.ExistByUserName(name)

	if check {
		return "", errors.New("user already exist")
	}

	user := repository.User{
		Username: name,
		Password: password,
	}

	//добавление и создание в мапу
	if err := s.userRepository.Create(&user); err != nil {
		return "", err
	}

	//Создание JWT токена

	jwtToken, err := jwt.GenerateJwtToken(name)
	if err != nil {
		return "", err
	}

	return jwtToken, nil

}

func (s *AuthService) LoginUser(name string, password string) (string, error) {

	//1. Проверить введены ли поля
	//2. Получить данные пользователя
	//3. Проверить пароль что он подходит
	//4. Если все корректно сгенерировать JWT токен и вернуть ответ

	if name == "" {
		return "", errors.New("username is required")
	}

	if password == "" {
		return "", errors.New("password is required")
	}

	//Получаем данные пользователя по username
	user, err := s.userRepository.FindByUsername(name)
	if err != nil {
		return "", err
	}

	//Проверка пароля
	if user.Password != password {
		return "", errors.New("Incorrect password")
	}

	jwtToken, err := jwt.GenerateJwtToken(user.Username)
	if err != nil {
		return "", err
	}

	return jwtToken, nil

}
