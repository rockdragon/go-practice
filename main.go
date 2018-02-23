package main

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/beego/mux"
	"fmt"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	serverAddr := "127.0.0.1:8080"
	fmt.Printf("server running upon: %s\n", serverAddr)

	mx := mux.New()
	mx.Handler("GET", "/", http.FileServer(http.Dir(".")))
	sugar.Fatal(http.ListenAndServe(serverAddr, mx))
}