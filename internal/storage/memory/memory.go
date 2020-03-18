package memory

import (
	"context"
	"errors"

	"github.com/fanadol/graphql-go-example/internal/user"
)

var users = []*user.User{
	&user.User{ID: 1, Name: "Eko", Friends: []int32{2}},
	&user.User{ID: 2, Name: "Echo", Friends: []int32{1, 3}},
	&user.User{ID: 3, Name: "Echa", Friends: []int32{1, 2}},
}

type Storage struct{}

// GetUsers is an in memory database for gRPC service
func (s *Storage) Get(ctx context.Context) ([]*user.User, error) {
	return users, nil
}

func (s *Storage) GetByUserID(ctx context.Context, ID int32) (*user.User, error) {
	for _, user := range users {
		if user.ID == ID {
			return user, nil
		}
	}

	return nil, errors.New("Not Found")
}
