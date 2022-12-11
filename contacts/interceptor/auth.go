package interceptor

import (
	"context"
	"github.com/iamyxsh/grpc-chat/contacts/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func CheckHeaderInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		headers, ok := metadata.FromIncomingContext(ctx)

		if !ok {
			return nil, status.Error(codes.Internal, "Error while reading the context")
		}

		token := headers.Get("auth")

		if len(token) == 0 {
			return nil, status.Error(codes.Unauthenticated, "Expected authorization header")
		}

		number, ok := utils.VerifyJWT(token[0])

		if !ok {
			return "", status.Error(
				codes.Unauthenticated,
				"unauthorized",
			)
		}

		headers.Append("number", number)
		ctx = metadata.NewIncomingContext(ctx, headers)

		return handler(ctx, req)
	}
}
