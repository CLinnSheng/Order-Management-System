package main

import (
	"log"
	"net/http"

	"github.com/CLinnSheng/OrderMangementSystem/common"
	_ "github.com/joho/godotenv/autoload"
)

var (
	httpAddr = common.GetEnvVar("HTTP_ADDR", ":3030")
)

// expose http server
// implement gRPC connection
func main() {

	// setting up http Server
	mux := http.NewServeMux()
	handler := NewHandler()

	handler.registerRoutes(mux)

	log.Printf("Starting HTTP server at %s", httpAddr)

	err := http.ListenAndServe(httpAddr, mux)
	mux.Hand
	if err != nil {
		log.Fatal("Fail to start http server")
	}
}
