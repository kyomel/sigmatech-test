package repository

import (
	"context"
	"log"
	"sigmatech-test/pkg/helper"
	"sigmatech-test/pkg/model"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type userRepo struct {
	db *sqlx.DB
}

type UserRepository interface {
	RegisterUser(ctx context.Context, req *model.RegisterUser) (*model.RegisterUserID, error)
	AddConsumer(ctx context.Context, req *model.RegisterUser, id int) error
	GetUserByEmail(ctx context.Context, email string) (bool, error)
	FindUserByEmail(ctx context.Context, email string) (*model.UserLogin, error)
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) RegisterUser(ctx context.Context, req *model.RegisterUser) (*model.RegisterUserID, error) {
	var result model.RegisterUserID

	tx, err := u.db.Begin()
	if err != nil {
		return nil, err
	}

	hashPassword, err := getHashPassword(req.Password)
	if err != nil {
		log.Println("Error on Hash Password => ", err)
		return nil, err
	}

	query := `
		INSERT INTO users(email, password, status, createdat)
		VALUES($1, $2, 'unactive', current_timestamp)
		RETURNING userid
	`

	row := tx.QueryRowContext(ctx, query, req.Email, hashPassword)
	err = row.Scan(&result.ID)
	if err != nil {
		log.Println("SQL error on RegisterUser => Execute Query and Scan", err)
		return nil, err
	}

	helper.CommitOrRollback(tx, err)

	return &result, nil
}

func (u *userRepo) AddConsumer(ctx context.Context, req *model.RegisterUser, id int) error {
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}

	query := `
		INSERT INTO consumers(fullname, userid)
		VALUES($1, $2)
	`

	_, err = tx.QueryContext(ctx, query, req.Fullname, id)
	if err != nil {
		log.Println("SQL error on AddConsumer => Execute Query", err)
		return err
	}

	helper.CommitOrRollback(tx, err)

	return err
}

func getHashPassword(password string) (string, error) {
	bytePassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (u *userRepo) GetUserByEmail(ctx context.Context, email string) (bool, error) {
	var result bool

	query := `
		SELECT EXISTS (SELECT 1 FROM users WHERE email = $1)
	`

	row := u.db.QueryRowContext(ctx, query, email)
	err := row.Scan(&result)
	if err != nil {
		log.Println("SQL error on GetUserByEmail => Execute Query and Scan", err)
		return false, err
	}

	return result, nil
}

func (u *userRepo) FindUserByEmail(ctx context.Context, email string) (*model.UserLogin, error) {
	var result model.UserLogin

	query := `
		SELECT email, password
		FROM users
		WHERE email = $1
	`

	row := u.db.QueryRowContext(ctx, query, email)
	err := row.Scan(&result.Email, &result.Password)
	if err != nil {
		log.Println("SQL error on FindUserByEmail => Execute Query and Scan", err)
		return nil, err
	}

	return &result, nil
}
