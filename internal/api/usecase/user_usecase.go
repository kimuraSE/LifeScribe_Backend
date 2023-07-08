package usecase

import (
	"LifeScribe_Backend/internal/api/model"
	"LifeScribe_Backend/internal/api/repository"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	bycrypt "golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	Login(user model.User) (string, error)
	SignUp(user model.User) (string, error)
	Destory(userId uint) error
}

type userUsecase struct {
	ur repository.IUserRepository
}

func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}

func (u *userUsecase) Login(user model.User) (string, error) {
	storeUser := model.User{}

	if err := u.ur.GetUserByEmail(&storeUser, user.Email); err != nil {
		return "", err
	}

	err := bycrypt.CompareHashAndPassword(storeUser.Password, user.Password)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storeUser.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (u *userUsecase) SignUp(user model.User) (string, error) {

	hash, err := bycrypt.GenerateFromPassword(user.Password, 10)
	if err != nil {
		return "", err
	}

	storeUser := model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: hash,
	}

	if err := u.ur.CreateUser(&storeUser); err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storeUser.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func (u *userUsecase) Destory(userId uint) error {
	if err := u.ur.DeleteUser(userId); err != nil {
		return err
	}
	return nil
}
