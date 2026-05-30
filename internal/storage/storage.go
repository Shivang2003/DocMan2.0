package storage

import (
	"github.com/Shivang2003/DocMan2.0/internal/types"
)

type UserStorage interface {
	CreateUser(user types.User) error
	// DeleteUser(id string) error
	// UpdateUser(id string) error
	// ReadUser(id string)
}