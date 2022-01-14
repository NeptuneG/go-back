package interceptors

import (
	"context"
	"errors"
	"reflect"

	"github.com/NeptuneG/go-back/pkg/auth"
	"github.com/NeptuneG/go-back/pkg/log"
	logField "github.com/NeptuneG/go-back/pkg/log/field"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func UnaryDefaultAuthInterceptor(methods ...string) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		log.Info("unary interceptor", logField.String("full_method", info.FullMethod))

		for _, method := range methods {
			if info.FullMethod == method {
				userID := reflect.Indirect(reflect.ValueOf(req)).FieldByName("UserId").String()
				token, err := getToken(ctx)
				if err != nil {
					return nil, status.Error(codes.Unauthenticated, err.Error())
				}
				if err := auth.Authorize(token, userID); err != nil {
					if err == auth.ErrInvalidToken {
						return nil, status.Error(codes.Unauthenticated, "invalid token")
					} else if err == auth.ErrNoPermission {
						return nil, status.Error(codes.PermissionDenied, "no permission")
					} else {
						return nil, status.Error(codes.Internal, "internal error")
					}
				} else {
					break
				}
			}
		}

		return handler(ctx, req)
	}
}

func getToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("authorization token is not provided")
	}
	values := md["authorization"]
	if len(values) == 0 {
		return "", errors.New("authorization token is not provided")
	}
	return values[0], nil
}
