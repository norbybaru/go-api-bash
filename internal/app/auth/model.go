package auth

import "time"

type User struct {
	Id        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Nickname  string    `json:"nickname,omitempty" db:"nickname"`
	Password  string    `json:"-" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (d *User) TableName() string {
	return "users"
}

func NewUser(name string, email string, password string, nickname string) *User {
	now := time.Now().UTC()
	return &User{
		Name:      name,
		Email:     email,
		Password:  password,
		Nickname:  nickname,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
