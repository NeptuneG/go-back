package scraper

import (
	"context"
	"os"
	"time"

	"github.com/NeptuneG/go-back/api/proto/scraper"
	"google.golang.org/grpc"
)

var scraper_service_host = os.Getenv("SCRAPER_SERVICE_HOST") + ":" + os.Getenv("SCRAPER_SERVICE_PORT")

func NewClient(ctx context.Context, dialOptions ...grpc.DialOption) (scraper.ScrapeServiceClient, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, scraper_service_host, dialOptions...)
	if err != nil {
		return nil, err
	}

	return scraper.NewScrapeServiceClient(conn), nil
}
