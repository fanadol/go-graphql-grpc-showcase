package user

import (
	"context"
)

type User struct {
	ID      int32   `json:"id"`
	Name    string  `json:"name"`
	Friends []int32 `json:"friends"`
}

type Service interface {
	Get(ctx context.Context) ([]*User, error)
	GetUserFriends(ctx context.Context, ID int32) ([]*User, error)
}

type Repository interface {
	Get(ctx context.Context) ([]*User, error)
	GetByUserID(ctx context.Context, ID int32) (*User, error)
}
