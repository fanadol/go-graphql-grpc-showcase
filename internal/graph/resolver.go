package graph

import userpb "github.com/fanadol/graphql-go-example/internal/user/proto"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	GRPCUserClient userpb.UsersClient
}
