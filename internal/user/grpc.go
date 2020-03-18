package user

import (
	"context"

	userpb "github.com/fanadol/graphql-go-example/internal/user/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type handler struct {
	service Service
}

func NewGRPCHandler(gserver *grpc.Server, service Service) {
	h := &handler{service: service}
	userpb.RegisterUsersServer(gserver, h)
}

func (h *handler) Get(ctx context.Context, empty *empty.Empty) (*userpb.GetResponse, error) {
	users, err := h.service.Get(ctx)
	if err != nil {
		return nil, err
	}

	var pbusers []*userpb.User
	for _, user := range users {
		pbusers = append(pbusers, &userpb.User{
			ID:   user.ID,
			Name: user.Name,
		})
	}

	return &userpb.GetResponse{
		Users: pbusers,
	}, nil
}

func (h *handler) GetUserFriends(ctx context.Context, req *userpb.GetRequest) (*userpb.GetUserFriendsResponse, error) {
	friends, err := h.service.GetUserFriends(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	var gfr []*userpb.User
	for _, friend := range friends {
		gfr = append(gfr, &userpb.User{
			ID:   friend.ID,
			Name: friend.Name,
		})
	}

	return &userpb.GetUserFriendsResponse{
		Users: gfr,
	}, nil
}
