package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/Ikedatomohiro/grpc-practice/infrastructure"
	"github.com/Ikedatomohiro/grpc-practice/utils"
)

func main() {
	if err := utils.LoadDotEnv(); err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}
	mux := infrastructure.NewServiceHandler()
	http.ListenAndServe(
		os.Getenv("HOST"),
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
