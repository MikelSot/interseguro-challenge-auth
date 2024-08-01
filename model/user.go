package model

import "time"

type User struct {
	ID        uint      `json:"id"`
	FirstName string    `json:"first_name"`
	Lastname  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u User) HasID() bool { return u.ID > 0 }
