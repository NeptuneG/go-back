package main

import (
	"context"
	"net/http"
	"os"

	auth "github.com/NeptuneG/go-back/api/proto/auth"
	live "github.com/NeptuneG/go-back/api/proto/live"
	payment "github.com/NeptuneG/go-back/api/proto/payment"
	scraper "github.com/NeptuneG/go-back/api/proto/scraper"
	"github.com/NeptuneG/go-back/pkg/log"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

var (
	auth_service_host    = os.Getenv("AUTH_SERVICE_HOST") + ":" + os.Getenv("AUTH_SERVICE_PORT")
	live_service_host    = os.Getenv("LIVE_SERVICE_HOST") + ":" + os.Getenv("LIVE_SERVICE_PORT")
	payment_service_host = os.Getenv("PAYMENT_SERVICE_HOST") + ":" + os.Getenv("PAYMENT_SERVICE_PORT")
	scraper_service_host = os.Getenv("SCRAPER_SERVICE_HOST") + ":" + os.Getenv("SCRAPER_SERVICE_PORT")
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

	if err := auth.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, auth_service_host, opts); err != nil {
		log.Error("Failed to register auth service", log.Field.Error(err))
		panic(err)
	}
	if err := live.RegisterLiveServiceHandlerFromEndpoint(ctx, mux, live_service_host, opts); err != nil {
		log.Error("Failed to register live service", log.Field.Error(err))
		panic(err)
	}
	if err := payment.RegisterPaymentServiceHandlerFromEndpoint(ctx, mux, payment_service_host, opts); err != nil {
		log.Error("Failed to register payment service", log.Field.Error(err))
		panic(err)
	}
	if err := scraper.RegisterScrapeServiceHandlerFromEndpoint(ctx, mux, scraper_service_host, opts); err != nil {
		log.Error("Failed to register scraper service", log.Field.Error(err))
		panic(err)
	}

	if err := http.ListenAndServe(":4000", mux); err != nil {
		log.Fatal("failed to start server", log.Field.Error(err))
		panic(err)
	}
}
