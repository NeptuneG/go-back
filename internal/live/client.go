package live

import (
	"context"
	"os"
	"time"

	"github.com/NeptuneG/go-back/api/proto/live"
	"google.golang.org/grpc"
)

var live_service_host = os.Getenv("LIVE_SERVICE_HOST") + ":" + os.Getenv("LIVE_SERVICE_PORT")

func NewClient(ctx context.Context, dialOptions ...grpc.DialOption) (live.LiveServiceClient, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, live_service_host, dialOptions...)
	if err != nil {
		return nil, err
	}

	return live.NewLiveServiceClient(conn), nil
}
