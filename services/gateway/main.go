package main

import (
	"context"
	"net/http"

	live "github.com/NeptuneG/go-back/gen/go/services/live/proto"
	payment "github.com/NeptuneG/go-back/gen/go/services/payment/proto"
	scraper "github.com/NeptuneG/go-back/gen/go/services/scraper/proto"
	user "github.com/NeptuneG/go-back/gen/go/services/user/proto"
	"github.com/NeptuneG/go-back/pkg/logger"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
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

	logger := logger.New()
	grpc_zap.ReplaceGrpcLogger(logger)

	server := http.Server{
		Addr:    ":4000",
		Handler: mux,
	}

	userConn, err := grpc.DialContext(ctx, "user-service:3377", opts...)
	if err != nil {
		logger.Fatal("failed to connect to user service", zap.Error(err))
		panic(err)
	}
	if err = user.RegisterUserServiceHandlerClient(ctx, mux, user.NewUserServiceClient(userConn)); err != nil {
		logger.Fatal("failed to register user service handler", zap.Error(err))
		panic(err)
	}

	liveConn, err := grpc.DialContext(ctx, "live-service:3377", opts...)
	if err != nil {
		logger.Fatal("failed to connect to live service", zap.Error(err))
		panic(err)
	}
	if err = live.RegisterLiveServiceHandlerClient(ctx, mux, live.NewLiveServiceClient(liveConn)); err != nil {
		logger.Fatal("failed to register live service handler", zap.Error(err))
		panic(err)
	}

	paymentConn, err := grpc.DialContext(ctx, "payment-service:3377", opts...)
	if err != nil {
		logger.Fatal("failed to connect to payment service", zap.Error(err))
		panic(err)
	}
	if err = payment.RegisterPaymentServiceHandlerClient(ctx, mux, payment.NewPaymentServiceClient(paymentConn)); err != nil {
		logger.Fatal("failed to register payment service handler", zap.Error(err))
		panic(err)
	}

	scraperConn, err := grpc.DialContext(ctx, "scraper-service:3377", opts...)
	if err != nil {
		logger.Fatal("failed to connect to scraper service", zap.Error(err))
		panic(err)
	}
	if err = scraper.RegisterScrapeServiceHandlerClient(ctx, mux, scraper.NewScrapeServiceClient(scraperConn)); err != nil {
		logger.Fatal("failed to register scraper service handler", zap.Error(err))
		panic(err)
	}

	if err := server.ListenAndServe(); err != nil {
		logger.Fatal("failed to start server", zap.Error(err))
		panic(err)
	}
}
