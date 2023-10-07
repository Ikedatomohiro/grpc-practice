package infrastructure

import (
	"net/http"

	todov1 "github.com/Ikedatomohiro/grpc-practice/todo/v1"
	"github.com/Ikedatomohiro/grpc-practice/todo/v1/todov1connect"

	userv1 "github.com/Ikedatomohiro/grpc-practice/user/v1"
	"github.com/Ikedatomohiro/grpc-practice/user/v1/userv1connect"
)

func NewServiceHandler() http.Handler {
	mux := http.NewServeMux()

	todoer := &todov1.TodoServer{}
	path, handler := todov1connect.NewTodoServiceHandler(todoer)
	mux.Handle(path, handler)

	userer := &userv1.UserServer{}
	path, handler = userv1connect.NewUserServiceHandler(userer)
	mux.Handle(path, handler)
	return mux
}
