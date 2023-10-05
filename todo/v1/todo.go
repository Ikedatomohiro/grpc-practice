package todov1

import (
	"context"
	"fmt"
	"log/slog"
	"sync"

	"connectrpc.com/connect"

	"github.com/google/uuid"
)

type TodoServer struct {
	items sync.Map
}

func (s *TodoServer) CreateTodo(
	ctx context.Context,
	req *connect.Request[CreateTodoRequest],
) (*connect.Response[CreateTodoResponse], error) {
	slog.Info("CreateTodo")
	id := uuid.New().String()
	item := &TodoItem{
		Id:     id,
		Title:  req.Msg.Title,
		Status: TodoItem_STATUS_OPEN,
	}
	s.items.Store(id, item)
	res := connect.NewResponse(&CreateTodoResponse{
		Item: item,
	})

	return res, nil
}

func (s *TodoServer) DeleteTodo(
	ctx context.Context,
	req *connect.Request[DeleteTodoRequest],
) (*connect.Response[DeleteTodoResponse], error) {
	slog.Info("DeleteTodo")
	_, ok := s.get(req.Msg.Id)
	if !ok {
		return nil, fmt.Errorf("item not found")
	}

	s.items.Delete(req.Msg.Id)
	res := connect.Response[DeleteTodoResponse]{
		Msg: &DeleteTodoResponse{
			Id: req.Msg.Id,
		},
	}
	return &res, nil
}

func (s *TodoServer) UpdateTodo(
	ctx context.Context,
	req *connect.Request[UpdateTodoRequest],
) (*connect.Response[UpdateTodoResponse], error) {
	slog.Info("UpdateTodo")
	todo, ok := s.get(req.Msg.Id)
	if !ok {
		return nil, fmt.Errorf("item not found")
	}
	todo.Status = req.Msg.Status
	s.items.Store(todo.Id, todo)
	res := connect.Response[UpdateTodoResponse]{
		Msg: &UpdateTodoResponse{
			Item: todo,
		},
	}
	return &res, nil
}

func (s *TodoServer) GetTodoList(
	ctx context.Context,
	req *connect.Request[GetTodoListRequest],
) (*connect.Response[GetTodoListResponse], error) {
	slog.Info("GetTodoList")
	var items []*TodoItem
	s.items.Range(func(key, value interface{}) bool {
		items = append(items, value.(*TodoItem))
		return true
	})
	res := connect.Response[GetTodoListResponse]{
		Msg: &GetTodoListResponse{
			Items: items,
		},
	}
	return &res, nil
}

func (s *TodoServer) get(id string) (*TodoItem, bool) {
	item, ok := s.items.Load(id)
	if ok {
		return item.(*TodoItem), true
	}
	return nil, false
}
