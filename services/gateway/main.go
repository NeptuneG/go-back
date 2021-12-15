package main

import (
	"context"
	"net/http"

	live "github.com/NeptuneG/go-back/services/live/proto"
	scraper "github.com/NeptuneG/go-back/services/scraper/proto"
	user "github.com/NeptuneG/go-back/services/user/proto"
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

	userConn, err := grpc.DialContext(ctx, "user-service:3377", opts...)
	if err != nil {
		panic(err)
	}
	if err = user.RegisterUserServiceHandlerClient(ctx, mux, user.NewUserServiceClient(userConn)); err != nil {
		panic(err)
	}

	liveConn, err := grpc.DialContext(ctx, "live-service:3377", opts...)
	if err != nil {
		panic(err)
	}
	if err = live.RegisterLiveServiceHandlerClient(ctx, mux, live.NewLiveServiceClient(liveConn)); err != nil {
		panic(err)
	}

	scraperConn, err := grpc.DialContext(ctx, "scraper-service:3377", opts...)
	if err != nil {
		panic(err)
	}
	if err = scraper.RegisterScrapeServiceHandlerClient(ctx, mux, scraper.NewScrapeServiceClient(scraperConn)); err != nil {
		panic(err)
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}