package user

import (
	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(DB *sqlx.DB) *userRepository {
	return &userRepository{DB: DB}
}
