package user

import "time"

type UserNoPassword struct {
	Uuid string `json:"uuid"`
	Name string `json:"name"`
	Email string `json:"email"`
	EmailVerified bool `json:"email_verified"`
	Active bool `json:"active"`
	CreatedAt time.Time `json:"created_at"`
}