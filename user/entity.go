package user

import "time"

// untuk mengakses database users di mysql
type User struct {
	ID             int
	Name           string
	Occupation     string
	Email		   string
	PasswordHash   string
	AvatarFileName string
	Role           string
	Token          string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}