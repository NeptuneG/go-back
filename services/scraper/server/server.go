package server

import (
	"context"
	"strings"

	"github.com/NeptuneG/go-back/gen/go/services/scraper/proto"
	faktory "github.com/contribsys/faktory/client"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ScrapeService struct {
	proto.UnimplementedScrapeServiceServer
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
