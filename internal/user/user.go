package user

import "context"

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) Get(ctx context.Context) ([]*User, error) {
	return s.repo.Get(ctx)
}

func (s *service) GetUserFriends(ctx context.Context, ID int32) ([]*User, error) {
	friends := make([]*User, 0)
	users, err := s.repo.Get(ctx)
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if ID == user.ID {
			for _, friend := range user.Friends {
				u, err := s.repo.GetByUserID(ctx, friend)
				if err != nil {
					return nil, err
				}

				friends = append(friends, u)
			}
		}
	}

	return friends, nil
}
