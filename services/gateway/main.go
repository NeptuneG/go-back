package main

import (
	"context"
	"net/http"

	liveProto "github.com/NeptuneG/go-back/services/live/proto"
	userProto "github.com/NeptuneG/go-back/services/user/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
	}

	server := http.Server{
		Addr:    ":4000",
		Handler: mux,
	}

	userConnection, err := grpc.DialContext(ctx, "user-service:3377", opts...)
	if err != nil {
		panic(err)
	}
	if err = userProto.RegisterUserServiceHandlerClient(ctx, mux, userProto.NewUserServiceClient(userConnection)); err != nil {
		panic(err)
	}

	liveConnection, err := grpc.DialContext(ctx, "live-service:3377", opts...)
	if err != nil {
		panic(err)
	}
	if err = liveProto.RegisterLiveServiceHandlerClient(ctx, mux, liveProto.NewLiveServiceClient(liveConnection)); err != nil {
		panic(err)
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
