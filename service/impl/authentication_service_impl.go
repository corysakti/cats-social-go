package impl

import (
	"context"
	"database/sql"
	"github.com/corysakti/cats-social-go/helper"
	"github.com/corysakti/cats-social-go/model/entity"
	"github.com/corysakti/cats-social-go/model/web/request"
	"github.com/corysakti/cats-social-go/model/web/response"
	"github.com/corysakti/cats-social-go/repository"
	"github.com/corysakti/cats-social-go/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthenticationServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validator      *validator.Validate
}

func NewAuthenticationService(userRepository repository.UserRepository, DB *sql.DB, validator *validator.Validate) service.AuthenticationService {
	return &AuthenticationServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validator:      validator}
}

func (service AuthenticationServiceImpl) Login(ctx context.Context, request request.LoginRequest) response.AuthenticationResponse {
	err := service.Validator.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	helper.PanicIfError(err)

	existUser, err := service.UserRepository.FindByEmail(ctx, tx, request.Email)
	helper.PanicThrowNotFound(err)

	err = bcrypt.CompareHashAndPassword([]byte(existUser.Password), []byte(request.Password))
	helper.PanicPasswordIsWrong(err)

	accessToken, err := createNewToken(existUser)
	helper.PanicIfError(err)

	return response.AuthenticationResponse{
		Email:       existUser.Email,
		Name:        existUser.Name,
		AccessToken: accessToken,
	}
}

func (service AuthenticationServiceImpl) Register(ctx context.Context, request request.RegisterRequest) response.AuthenticationResponse {
	err := service.Validator.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, err = service.UserRepository.FindByEmail(ctx, tx, request.Email)
	helper.PanicThrowAlreadyExist(err)

	has, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	helper.PanicIfError(err)

	user := entity.User{
		Email:    request.Email,
		Name:     request.Name,
		Password: string(has),
	}

	user = service.UserRepository.Save(ctx, tx, user)

	accessToken, err := createNewToken(user)
	helper.PanicIfError(err)

	return response.AuthenticationResponse{
		Email:       user.Email,
		Name:        user.Name,
		AccessToken: accessToken,
	}
}

func createNewToken(user entity.User) (string, error) {
	// Define the expiration time for the token (current time + 8 hours)
	expirationTime := time.Now().Add(8 * time.Hour)

	// Create the claims for the token
	claims := jwt.MapClaims{
		"id":    user.Id,
		"email": user.Email,
		"name":  user.Name,
		"exp":   expirationTime.Unix(), // Expiration time in Unix time
	}

	// Create the token with the claims and the signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a secret key
	accessToken, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
