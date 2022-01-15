package interceptors

import (
	"context"
	"errors"
	"reflect"

	"github.com/NeptuneG/go-back/pkg/auth"
	"github.com/NeptuneG/go-back/pkg/log"

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
		log.Info("unary interceptor", log.Field.String("full_method", info.FullMethod))

		if isToVerify(info.FullMethod, methods...) {
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
			}
		}

		return handler(ctx, req)
	}
}

func isToVerify(fullMethod string, methods ...string) bool {
	if methods[0] == "*" {
		return true
	}
	for _, method := range methods {
		if fullMethod == method {
			return true
		}
	}
	return false
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
