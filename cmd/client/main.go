package main

import (
	"context"
	"log"
	"net/http"

	todov1 "example.com/todo/gen/todo/v1"
	"example.com/todo/gen/todo/v1/todov1connect"

	"connectrpc.com/connect"
)

func main() {
	client := todov1connect.NewTodoServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
	)
	res, err := client.CreateTodo(
		context.Background(),
		connect.NewRequest(&todov1.CreateTodoRequest{Title: "create todo"}),
	)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(res.Msg.Item)
}
