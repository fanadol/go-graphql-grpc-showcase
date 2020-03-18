package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/fanadol/graphql-go-example/internal/graph/generated"
	"github.com/fanadol/graphql-go-example/internal/graph/model"
	userpb "github.com/fanadol/graphql-go-example/internal/user/proto"
	"github.com/golang/protobuf/ptypes/empty"
)

func (r *queryResolver) User(ctx context.Context, input int32) (*model.User, error) {
	userList, err := r.GRPCUserClient.Get(ctx, &empty.Empty{})
	if err != nil {
		return nil, err
	}

	users := userList.Users
	for _, user := range users {
		if user.ID == input {
			return &model.User{
				ID:   user.ID,
				Name: user.Name,
			}, nil
		}
	}

	return nil, graphql.DefaultErrorPresenter(ctx, errors.New("Not Found"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	userList, err := r.GRPCUserClient.Get(ctx, &empty.Empty{})
	if err != nil {
		return nil, err
	}

	var users []*model.User
	for _, user := range userList.Users {
		users = append(users, &model.User{
			ID:   user.ID,
			Name: user.Name,
		})
	}

	return users, nil
}

func (r *userResolver) Friends(ctx context.Context, obj *model.User) ([]*model.User, error) {
	friends := make([]*model.User, 0)

	friendList, err := r.GRPCUserClient.GetUserFriends(ctx, &userpb.GetRequest{ID: obj.ID})
	if err != nil {
		return nil, err
	}

	for _, friend := range friendList.Users {
		friends = append(friends, &model.User{
			ID:   friend.ID,
			Name: friend.Name,
		})
	}

	return friends, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
