package user

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type userRepository struct {
	DB *sqlx.DB
}

type User struct {
	Id        string
	Name      string
	Token     string
	BlockedAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUserRepository(DB *sqlx.DB) *userRepository {
	return &userRepository{DB: DB}
}
