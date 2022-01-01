package main

import (
	"context"
	"net/http"

	live "github.com/NeptuneG/go-back/gen/go/services/live/proto"
	payment "github.com/NeptuneG/go-back/gen/go/services/payment/proto"
	scraper "github.com/NeptuneG/go-back/gen/go/services/scraper/proto"
	user "github.com/NeptuneG/go-back/gen/go/services/user/proto"
	"github.com/NeptuneG/go-back/pkg/log"
	logField "github.com/NeptuneG/go-back/pkg/log/field"
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

	userConn, err := grpc.DialContext(ctx, "user.default.svc.cluster.local:3377", opts...)
	if err != nil {
		log.Fatal("failed to connect to user service", logField.Error(err))
		panic(err)
	}
	if err = user.RegisterUserServiceHandlerClient(ctx, mux, user.NewUserServiceClient(userConn)); err != nil {
		log.Fatal("failed to register user service handler", logField.Error(err))
		panic(err)
	}

	liveConn, err := grpc.DialContext(ctx, "live.default.svc.cluster.local:3377", opts...)
	if err != nil {
		log.Fatal("failed to connect to live service", logField.Error(err))
		panic(err)
	}
	if err = live.RegisterLiveServiceHandlerClient(ctx, mux, live.NewLiveServiceClient(liveConn)); err != nil {
		log.Fatal("failed to register live service handler", logField.Error(err))
		panic(err)
	}

	paymentConn, err := grpc.DialContext(ctx, "payment.default.svc.cluster.local:3377", opts...)
	if err != nil {
		log.Fatal("failed to connect to payment service", logField.Error(err))
		panic(err)
	}
	if err = payment.RegisterPaymentServiceHandlerClient(ctx, mux, payment.NewPaymentServiceClient(paymentConn)); err != nil {
		log.Fatal("failed to register payment service handler", logField.Error(err))
		panic(err)
	}

	scraperConn, err := grpc.DialContext(ctx, "scraper.default.svc.cluster.local:3377", opts...)
	if err != nil {
		log.Fatal("failed to connect to scraper service", logField.Error(err))
		panic(err)
	}
	if err = scraper.RegisterScrapeServiceHandlerClient(ctx, mux, scraper.NewScrapeServiceClient(scraperConn)); err != nil {
		log.Fatal("failed to register scraper service handler", logField.Error(err))
		panic(err)
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("failed to start server", logField.Error(err))
		panic(err)
	}
}
