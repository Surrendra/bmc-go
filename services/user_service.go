package services

import (
	"BaliMediaCenter/middlewares"
	"BaliMediaCenter/models"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
}

func NewUserService() *userService {
	return &userService{}
}

type UserService interface {
	Login(username string, password string) (interface{}, error, string)
}

func (s userService) Login(username string, password string) (interface{}, error, string) {
	user := models.User{}
	models.DB.Where("username = ?", username).First(&user)
	// if user is empty
	if user.Username != username {
		return nil, errors.New("username not match"), "username not found in our database"
	}
	errHash := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errHash != nil {
		return nil, errHash, "Username or Password is wrong"
	}
	token, errGenToken := middlewares.GenerateJWT(user.Code)
	if errGenToken != nil {
		return nil, errHash, "Gagal ketika proses generate Token!"
	}
	user.LastToken = token
	return user, nil, "Success"
}

func Logging() {
	fmt.Println("log with goroutine")
}
