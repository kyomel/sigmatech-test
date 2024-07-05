package usecase

import (
	"context"
	"errors"
	"log"
	"sigmatech-test/pkg/libs/middleware"
	"sigmatech-test/pkg/model"
	"sigmatech-test/pkg/repository"
	"time"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	contextTimeout time.Duration
	userRepository repository.UserRepository
	db             *sqlx.DB
}

type UserUseCase interface {
	RegisterUser(ctx context.Context, request *model.RegisterUser) error
	Login(ctx context.Context, request *model.LoginParam) (*model.TokenLogin, error)
}

func NewUserUseCase(contextTimeout time.Duration, userRepository repository.UserRepository, db *sqlx.DB) UserUseCase {
	return &userUsecase{
		contextTimeout,
		userRepository,
		db,
	}
}

func (u *userUsecase) RegisterUser(ctx context.Context, request *model.RegisterUser) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	checkUser, err := u.userRepository.GetUserByEmail(ctx, request.Email)
	if err != nil {
		log.Println("Error in UseCase RegisterUser => Open Transaction GetUserByEmail => ", err)
		return err
	}

	if checkUser {
		return errors.New("email already exist")
	}

	res, err := u.userRepository.RegisterUser(ctx, request)
	if err != nil {
		log.Println("Error in UseCase RegisterUser => Open Transaction RegisterUser => ", err)
		return err
	}

	err = u.userRepository.AddConsumer(ctx, request, res.ID)
	if err != nil {
		log.Println("Error in UseCase RegisterUser => Open Transaction AddConsumer => ", err)
		return err
	}

	return nil
}

func (u *userUsecase) Login(ctx context.Context, request *model.LoginParam) (*model.TokenLogin, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var result *model.TokenLogin

	user, err := u.userRepository.FindUserByEmail(ctx, request.Email)
	if err != nil {
		log.Println("Error in UseCase Login => Open Transaction FindUserByEmail => ", err)
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		log.Println("Error in UseCase Login => Open Transaction CompareHashAndPassword => ", err)
		return nil, err
	}

	token, err := middleware.GenerateToken(user.Email)
	if err != nil {
		log.Println("Error in UseCase Login => Open Transaction GenerateToken => ", err)
		return nil, err
	}

	result = &model.TokenLogin{
		TokenJWT: token,
	}

	return result, nil
}
