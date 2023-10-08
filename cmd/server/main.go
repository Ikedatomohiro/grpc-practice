package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/Ikedatomohiro/grpc-practice/infrastructure"
)

func main() {
	fmt.Println("Server start")
	mux := infrastructure.NewServiceHandler()
	http.ListenAndServe(
		"localhost:8080",
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
