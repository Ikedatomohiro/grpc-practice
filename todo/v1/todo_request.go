package todov1

import (
	"fmt"

	"connectrpc.com/connect"
)

func validate(r *connect.Request[CreateTodoRequest]) error {
	if r.Msg.Title == "" {
		return fmt.Errorf("title is empty")
	}
	if r.Msg.UserId == "" {
		return fmt.Errorf("user id is empty")
	}
	return nil
}
