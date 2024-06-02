package services

import (
	"gin-fleamarket/models"
	"gin-fleamarket/repositories"

	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Signup(email string, password string) error
}

type AuthService struct {
	repository repositories.IAuthRepository
}

func NewAuthService(repository repositories.IAuthRepository) IAuthService {
	return &AuthService{repository: repository}
}

func (s *AuthService) Signup(email string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //文字列はイミュータブルで、バイトスライスに変換することで、ミュータブル化。また、ハッシュ関数はバイナリデータを処理するため、パスワードをバイトスライスに変換することで、処理できる
	//コストパラメータは、デフォルトだと10(範囲4～31)
	if err != nil {
		return err
	}
	user := models.User{
		Email: email,
		Password: string(hashedPassword),
	}
	return s.repository.CreateUser(user)
}