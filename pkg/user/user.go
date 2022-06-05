package user

import "time"

type User struct {
	Id        string
	Name      string
	Token     string
	BlockedAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
