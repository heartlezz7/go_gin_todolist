package model

import "time"

type UserRole struct {
	User  string
	Admin string
}

var UserRoles = UserRole{
	User:  "user",
	Admin: "admin",
}

type UserData struct {
	Id           int       `db:"id"`
	Username     string    `db:"username"`
	Email        string    `db:"email"`
	Password     string    `db:"password"`
	ProfileImage *string   `db:"profileImage"`
	Role         string    `db:"role"`
	Enable       bool      `db:"enable"`
	CreatedAt    time.Time `db:"created_at"`
}

type CreateUser struct {
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	ProfileImage *string   `json:"profileImage"`
	Role         *UserRole `json:"role"`
	Enable       *bool     `json:"enable"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
