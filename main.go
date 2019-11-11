package main

import (
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	echopb "github.com/microservicedemo/echo/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	gmux := runtime.NewServeMux()
	ctx := context.Background()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := echopb.RegisterEchoServiceHandlerFromEndpoint(ctx, gmux, "echo:9090", opts)
	if err != nil {
		log.Fatal(err)
		return
	}

	httpmux := http.NewServeMux()
	httpmux.Handle("/", gmux)
	http.ListenAndServe(":9091", httpmux)
}
