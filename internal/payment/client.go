package payment

import (
	"context"
	"os"
	"time"

	"github.com/NeptuneG/go-back/api/proto/payment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var payment_service_host = os.Getenv("PAYMENT_SERVICE_HOST") + ":" + os.Getenv("PAYMENT_SERVICE_PORT")

func NewClient() (payment.PaymentServiceClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, payment_service_host, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	return payment.NewPaymentServiceClient(conn), nil
}
