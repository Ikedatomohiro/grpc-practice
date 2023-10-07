package userv1

import (
	"context"
	"fmt"
	"sync"

	"log/slog"

	"connectrpc.com/connect"
	"github.com/google/uuid"
)

type UserServer struct {
	items sync.Map
}

func (s *UserServer) CreateUser(ctx context.Context,
	req *connect.Request[CreateUserRequest],
) (*connect.Response[CreateUserResponse], error) {
	slog.Info("CreateUser")
	id := uuid.New().String()
	user := &User{
		Id:    id,
		Name:  req.Msg.Name,
		Email: req.Msg.Email,
	}
	s.items.Store(id, user)
	res := connect.NewResponse(&CreateUserResponse{
		User: user,
	})

	return res, nil
}

func (s *UserServer) GetUser(ctx context.Context,
	req *connect.Request[GetUserRequest],
) (*connect.Response[GetUserResponse], error) {
	slog.Info("GetUser")
	user, ok := s.get(req.Msg.Id)
	if !ok {
		return nil, fmt.Errorf("user not found")
	}

	res := connect.NewResponse(&GetUserResponse{
		User: user,
	})

	return res, nil
}

func (s *UserServer) get(id string) (*User, bool) {
	v, ok := s.items.Load(id)
	if !ok {
		return nil, false
	}
	user, ok := v.(*User)
	if !ok {
		return nil, false
	}
	return user, true
}
