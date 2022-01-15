package live

import (
	"context"
	"os"
	"time"

	"github.com/NeptuneG/go-back/api/proto/live"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var live_service_host = os.Getenv("LIVE_SERVICE_HOST") + ":" + os.Getenv("LIVE_SERVICE_PORT")

func NewClient() (live.LiveServiceClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, live_service_host, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	return live.NewLiveServiceClient(conn), nil
}
