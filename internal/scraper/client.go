package scraper

import (
	"context"
	"os"
	"time"

	"github.com/NeptuneG/go-back/api/proto/scraper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var scraper_service_host = os.Getenv("SCRAPER_SERVICE_HOST") + ":" + os.Getenv("SCRAPER_SERVICE_PORT")

func NewClient() (scraper.ScrapeServiceClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, scraper_service_host, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	return scraper.NewScrapeServiceClient(conn), nil
}