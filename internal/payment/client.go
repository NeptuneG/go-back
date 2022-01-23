package payment

import (
	"context"
	"os"
	"time"

	"github.com/NeptuneG/go-back/api/proto/payment"
	"google.golang.org/grpc"
)

var payment_service_host = os.Getenv("PAYMENT_SERVICE_HOST") + ":" + os.Getenv("PAYMENT_SERVICE_PORT")

func NewClient(ctx context.Context, dialOptions ...grpc.DialOption) (payment.PaymentServiceClient, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, payment_service_host, dialOptions...)
	if err != nil {
		return nil, err
	}

	return payment.NewPaymentServiceClient(conn), nil
}
