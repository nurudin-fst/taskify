package usecase

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nurudin-fst/taskify/internal/dto"
	"github.com/nurudin-fst/taskify/internal/helper"
	"github.com/nurudin-fst/taskify/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserUC struct {
	UserRepo repository.UserRepository
}

func NewUserUC(userRepo *repository.UserRepository) *UserUC {
	return &UserUC{
		UserRepo: *userRepo,
	}
}

func (uc *UserUC) Register(in dto.UserRegisterIn) (code int, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return fiber.StatusInternalServerError, err
	}
	userInsert := repository.User{
		Name:     in.Name,
		Email:    in.Email,
		Password: string(hashedPassword),
	}
	err = uc.UserRepo.Insert(userInsert)
	if err != nil {
		return fiber.StatusInternalServerError, err
	}

	return fiber.StatusCreated, err
}

func (uc *UserUC) Login(in dto.UserLoginIn) (token string, code int, err error) {
	user, err := uc.UserRepo.GetUserByEmail(in.Email)
	if err != nil {
		return "", fiber.StatusUnauthorized, err
	}
	if !helper.IsValidPassword(in.Password, user.Password) {
		return "", fiber.StatusUnauthorized, errors.New("email or password is invalid")
	}

	tokenExpire := time.Now().Add(48 * time.Hour)
	token, err = helper.GenerateJWT(user.Id, tokenExpire)
	if err != nil {
		return "", fiber.StatusInternalServerError, err
	}

	return token, fiber.StatusOK, err
}
