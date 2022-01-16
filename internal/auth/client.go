package auth

import (
	"context"
	"os"
	"time"

	"github.com/NeptuneG/go-back/api/proto/auth"
	"google.golang.org/grpc"
)

var auth_service_host = os.Getenv("AUTH_SERVICE_HOST") + ":" + os.Getenv("AUTH_SERVICE_PORT")

func NewClient(dialOptions ...grpc.DialOption) (auth.AuthServiceClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, auth_service_host, dialOptions...)
	if err != nil {
		return nil, err
	}

	return auth.NewAuthServiceClient(conn), nil
}
