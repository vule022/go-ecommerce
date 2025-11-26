package service

import (
	"go-microservices/internal/domain"
	"log"
)

type UserService struct {
}

func (s UserService) findUserByEmail(email string) (*domain.User, error) {
	//logic
	return nil, nil
}

func (s UserService) Register(input interface{}) (string, error) {
	log.Println(input)
	//logic
	return "register-token", nil
}

func (s UserService) Login(input interface{}) (string, error) {
	//logic
	return "", nil
}

func (s UserService) GetVerificationCode(e domain.User) (int, error) {
	//logic
	return 0, nil
}

func (s UserService) VerifyCode(id uint, code int) error {
	//logic
	return nil
}

func (s UserService) CreateProfile(id uint, input any) error {
	//logic
	return nil
}

func (s UserService) GetProfile(id uint) (*domain.User, error) {
	//logic
	return nil, nil
}

func (s UserService) UpdateProfile(id uint, input any) error {
	//logic
	return nil
}

func (s UserService) BecomeSeller(id uint, input any) (string, error) {
	//logic
	return "", nil
}

func (s UserService) FindCart(id uint) ([]interface{}, error) {
	//logic
	return nil, nil
}

func (s UserService) CreateCart(input any, u domain.User) ([]interface{}, error) {
	//logic
	return nil, nil
}

func (s UserService) CreateOrder(u domain.User) (int, error) {
	//logic
	return 0, nil
}

func (s UserService) GetOrders(u domain.User) ([]interface{}, error) {
	//logic
	return nil, nil
}

func (s UserService) GetOrderById(id uint, uId uint) (interface{}, error) {
	//logic
	return nil, nil
}
