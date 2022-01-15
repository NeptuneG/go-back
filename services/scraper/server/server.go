package server

import (
	"context"
	"os"
	"strings"

	liveProto "github.com/NeptuneG/go-back/api/proto/live"
	proto "github.com/NeptuneG/go-back/api/proto/scraper"
	"github.com/NeptuneG/go-back/pkg/log"
	"github.com/NeptuneG/go-back/services/scraper/consumer"
	faktory "github.com/contribsys/faktory/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ScrapeService struct {
	proto.UnimplementedScrapeServiceServer
	liveClient liveProto.LiveServiceClient
	consumer   *consumer.ScrapedEventsConsumer
}

func New() *ScrapeService {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
	}
	conn, err := grpc.DialContext(context.Background(), os.Getenv("LIVE_SERVICE_HOST")+":"+os.Getenv("LIVE_SERVICE_PORT"), opts...)
	if err != nil {
		log.Fatal("failed to dial live service", log.Field.Error(err))
		panic(err)
	}

	liveClient := liveProto.NewLiveServiceClient(conn)
	consumer := consumer.New(liveClient)
	consumer.Start()
	return &ScrapeService{liveClient: liveClient, consumer: consumer}
}

func (s *ScrapeService) Close() {
	s.consumer.Close()
}

func (s *ScrapeService) CreateScrapeLiveEventsJob(ctx context.Context, req *proto.CreateScrapeLiveEventsJobRequest) (*proto.CreateScrapeLiveEventsJobResponse, error) {
	client, err := faktory.Open()
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to connect to faktory")
	}

	// TODO: year_month validation
	job := faktory.NewJob(slugToJobName(req.LiveHouseSlug), req.YearMonth)
	if err = client.Push(job); err != nil {
		return nil, status.Error(codes.Internal, "failed to push job to faktory")
	}

	return &proto.CreateScrapeLiveEventsJobResponse{
		JobId: job.Jid,
	}, nil
}

func slugToJobName(slug string) string {
	words := strings.Split(slug, "-")
	for i, word := range words {
		words[i] = strings.Title(word)
	}
	return "Scrape" + strings.Join(words, "") + "Job"
}
