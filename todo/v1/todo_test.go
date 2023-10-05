package todov1_test

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	todov1 "github.com/Ikedatomohiro/grpc-practice/todo/v1"
)

func TestCreateTodo(t *testing.T) {
	var tests = []struct {
		name       string
		req        *connect.Request[todov1.CreateTodoRequest]
		wantTitle  string
		wantStatus todov1.TodoItem_Status
		wantErr    bool
	}{
		{
			name: "success: create todo",
			req: &connect.Request[todov1.CreateTodoRequest]{
				Msg: &todov1.CreateTodoRequest{
					Title: "test",
				},
			},
			wantTitle:  "test",
			wantStatus: todov1.TodoItem_STATUS_OPEN,
			wantErr:    false,
		},
	}
	s := &todov1.TodoServer{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.CreateTodo(context.Background(), tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoServer.CreateTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got.Msg.Item.Title != tt.wantTitle {
				t.Errorf("TodoServer.CreateTodo() got = %v, want %v", got.Msg.Item.Title, tt.wantTitle)
			}

			if got.Msg.Item.Status != tt.wantStatus {
				t.Errorf("TodoServer.CreateTodo() got = %v, want %v", got.Msg.Item.Status, tt.wantStatus)
			}
		})
	}
}
