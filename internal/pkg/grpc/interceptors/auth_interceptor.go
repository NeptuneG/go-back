package interceptors

import (
	"context"
	"errors"
	"reflect"
	"strings"

	"github.com/NeptuneG/go-back/internal/pkg/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func UnaryDefaultAuthInterceptor(methods ...string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if isToVerify(info.FullMethod, methods...) {
			token, err := getToken(ctx)
			if err != nil {
				return nil, status.Error(codes.Unauthenticated, err.Error())
			}
			userCliams, err := auth.Authenticate(token)
			if err != nil {
				return nil, status.Error(codes.Unauthenticated, err.Error())
			}
			userID := reflect.Indirect(reflect.ValueOf(req)).FieldByName("UserId").String()
			if err := auth.AuthorizeByUserID(userCliams, userID); err != nil {
				return nil, status.Error(codes.PermissionDenied, err.Error())
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
		return "", errors.New("metadata not found")
	}
	values := md["authorization"]
	if len(values) == 0 {
		return "", errors.New("authorization token is not provided")
	}
	authorization := strings.SplitN(values[0], " ", 2)
	if len(authorization) != 2 || !strings.EqualFold(authorization[0], "Bearer") {
		return "", errors.New("bad authorization string")
	}
	return authorization[1], nil
}
