package main

import (
	"context"
	"log"
	"net/http"

	live "github.com/NeptuneG/go-back/gen/go/services/live/proto"
	payment "github.com/NeptuneG/go-back/gen/go/services/payment/proto"
	scraper "github.com/NeptuneG/go-back/gen/go/services/scraper/proto"
	user "github.com/NeptuneG/go-back/gen/go/services/user/proto"
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
		log.Fatalln("failed to dial user service:", err)
		panic(err)
	}
	if err = user.RegisterUserServiceHandlerClient(ctx, mux, user.NewUserServiceClient(userConn)); err != nil {
		log.Fatalln("failed to register user service handler:", err)
		panic(err)
	}

	liveConn, err := grpc.DialContext(ctx, "live-service:3377", opts...)
	if err != nil {
		log.Fatalln("failed to dial live service:", err)
		panic(err)
	}
	if err = live.RegisterLiveServiceHandlerClient(ctx, mux, live.NewLiveServiceClient(liveConn)); err != nil {
		log.Fatalln("failed to register live service handler:", err)
		panic(err)
	}

	paymentConn, err := grpc.DialContext(ctx, "payment-service:3377", opts...)
	if err != nil {
		log.Fatalln("failed to dial payment service:", err)
		panic(err)
	}
	if err = payment.RegisterPaymentServiceHandlerClient(ctx, mux, payment.NewPaymentServiceClient(paymentConn)); err != nil {
		log.Fatalln("failed to register payment service handler:", err)
		panic(err)
	}

	scraperConn, err := grpc.DialContext(ctx, "scraper-service:3377", opts...)
	if err != nil {
		log.Fatalln("failed to dial scraper service:", err)
		panic(err)
	}
	if err = scraper.RegisterScrapeServiceHandlerClient(ctx, mux, scraper.NewScrapeServiceClient(scraperConn)); err != nil {
		log.Fatalln("failed to register scraper service handler:", err)
		panic(err)
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalln("failed to listen and serve:", err)
		panic(err)
	}
}
