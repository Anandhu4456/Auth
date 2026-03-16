package domain

import "time"

type User struct {
	ID          int64
	Name        string
	Username    string
	Email       string
	Password    string
	PhoneNumber string
	Address     string
	CreatedAt   time.Time
}
