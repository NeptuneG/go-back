package server

import (
	"context"

	"github.com/NeptuneG/go-back/api/proto/auth"
	"github.com/NeptuneG/go-back/api/proto/gateway"
	"github.com/NeptuneG/go-back/api/proto/live"
	"github.com/NeptuneG/go-back/api/proto/payment"
	"github.com/NeptuneG/go-back/api/proto/scraper"
	authSvc "github.com/NeptuneG/go-back/internal/auth"
	liveSvc "github.com/NeptuneG/go-back/internal/live"
	paymentSvc "github.com/NeptuneG/go-back/internal/payment"
	scraperSvc "github.com/NeptuneG/go-back/internal/scraper"
)

var prefix = "/" + gateway.GatewayService_ServiceDesc.ServiceName + "/"

var AuthRequiredMethods = []string{
	prefix + "GetUserPoints",
	prefix + "CreateUserPoints",
	prefix + "CreateLiveEventOrder",
}

type GatewayServer struct {
	gateway.UnimplementedGatewayServiceServer
	authClient    auth.AuthServiceClient
	liveClient    live.LiveServiceClient
	paymentClient payment.PaymentServiceClient
	scraperClient scraper.ScrapeServiceClient
}

func New() *GatewayServer {
	authClient := make(chan auth.AuthServiceClient)
	liveClient := make(chan live.LiveServiceClient)
	paymentClient := make(chan payment.PaymentServiceClient)
	scraperClient := make(chan scraper.ScrapeServiceClient)
	go func() {
		client, err := authSvc.NewClient()
		if err != nil {
			panic(err)
		}
		authClient <- client
	}()
	go func() {
		client, err := liveSvc.NewClient()
		if err != nil {
			panic(err)
		}
		liveClient <- client
	}()
	go func() {
		client, err := paymentSvc.NewClient()
		if err != nil {
			panic(err)
		}
		paymentClient <- client
	}()
	go func() {
		client, err := scraperSvc.NewClient()
		if err != nil {
			panic(err)
		}
		scraperClient <- client
	}()
	return &GatewayServer{
		authClient:    <-authClient,
		liveClient:    <-liveClient,
		paymentClient: <-paymentClient,
		scraperClient: <-scraperClient,
	}
}

func (s *GatewayServer) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	return s.authClient.Register(ctx, req)
}

func (s *GatewayServer) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	return s.authClient.Login(ctx, req)
}

func (s *GatewayServer) GetUserPoints(ctx context.Context, req *payment.GetUserPointsRequest) (*payment.GetUserPointsResponse, error) {
	return s.paymentClient.GetUserPoints(ctx, req)
}

func (s *GatewayServer) CreateUserPoints(ctx context.Context, req *payment.CreateUserPointsRequest) (*payment.CreateUserPointsResponse, error) {
	return s.paymentClient.CreateUserPoints(ctx, req)
}

func (s *GatewayServer) ListLiveHouses(ctx context.Context, req *live.ListLiveHousesRequest) (*live.ListLiveHousesResponse, error) {
	return s.liveClient.ListLiveHouses(ctx, req)
}

func (s *GatewayServer) CreateLiveEventOrder(ctx context.Context, req *payment.CreateLiveEventOrderRequest) (*payment.CreateLiveEventOrderResponse, error) {
	return s.paymentClient.CreateLiveEventOrder(ctx, req)
}

func (s *GatewayServer) CreateScrapeLiveEventsJob(ctx context.Context, req *scraper.CreateScrapeLiveEventsJobRequest) (*scraper.CreateScrapeLiveEventsJobResponse, error) {
	return s.scraperClient.CreateScrapeLiveEventsJob(ctx, req)
}
