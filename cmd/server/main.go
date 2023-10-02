package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	todov1 "example.com/todo/gen/todo/v1"        // generated by protoc-gen-go
	"example.com/todo/gen/todo/v1/todov1connect" // generated by protoc-gen-connect-go
	"github.com/google/uuid"
)

type TodoServer struct {
	items sync.Map
}

func (s *TodoServer) CreateTodo(
	ctx context.Context,
	req *connect.Request[todov1.CreateTodoRequest],
) (*connect.Response[todov1.CreateTodoResponse], error) {
	id := uuid.New().String()
	item := &todov1.TodoItem{
		Id:     id,
		Title:  req.Msg.Title,
		Status: todov1.TodoItem_STATUS_OPEN,
	}
	s.items.Store(id, item)
	res := connect.NewResponse(&todov1.CreateTodoResponse{
		Item: item,
	})
	fmt.Println("CreateTodo")

	return res, nil
}

func (s *TodoServer) DeleteTodo(
	ctx context.Context,
	req *connect.Request[todov1.DeleteTodoRequest],
) (*connect.Response[todov1.DeleteTodoResponse], error) {
	fmt.Println("DeleteTodo")
	_, ok := s.get(req.Msg.Id)
	if !ok {
		return nil, fmt.Errorf("item not found")
	}

	s.items.Delete(req.Msg.Id)
	res := connect.Response[todov1.DeleteTodoResponse]{
		Msg: &todov1.DeleteTodoResponse{
			Id: req.Msg.Id,
		},
	}
	return &res, nil
}

func (s *TodoServer) UpdateTodo(
	ctx context.Context,
	req *connect.Request[todov1.UpdateTodoRequest],
) (*connect.Response[todov1.UpdateTodoResponse], error) {
	fmt.Println("UpdateTodo")
	todo, ok := s.get(req.Msg.Id)
	if !ok {
		return nil, fmt.Errorf("item not found")
	}
	todo.Status = req.Msg.Status
	s.items.Store(todo.Id, todo)
	res := connect.Response[todov1.UpdateTodoResponse]{
		Msg: &todov1.UpdateTodoResponse{
			Item: todo,
		},
	}
	return &res, nil
}

func main() {
	todoer := &TodoServer{}
	mux := http.NewServeMux()
	path, handler := todov1connect.NewTodoServiceHandler(todoer)
	mux.Handle(path, handler)
	http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}

func (s *TodoServer) get(id string) (*todov1.TodoItem, bool) {
	item, ok := s.items.Load(id)
	if ok {
		return item.(*todov1.TodoItem), true
	}
	return nil, false
}
