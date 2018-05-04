package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/golang/glog"

	gw "github.com/alextanhongpin/go-video/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

var videoEndpoint = flag.String("video_endpoint", "localhost:9090", "endpoint of video service")

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterVideoServiceHandlerFromEndpoint(ctx, mux, *videoEndpoint, opts)
	if err != nil {
		return err
	}

	srv := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      mux,
	}

	fmt.Println("listening to port *:8080. press ctrl + c to cancel.")
	return srv.ListenAndServe()
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
