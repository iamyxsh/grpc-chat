package interceptor

import (
	"context"

	"github.com/iamyxsh/grpc-chat/chat/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type serverStream struct {
	grpc.ServerStream
	ctx context.Context
}

func StreamAuthInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		headers, ok := metadata.FromIncomingContext(stream.Context())
		if !ok {
			return status.Error(codes.Internal, "Error while reading the context")
		}
		token := headers.Get("auth")
		if len(token) == 0 {
			return status.Error(codes.Unauthenticated, "Expected authorization header")
		}
		number, ok := utils.VerifyJWT(token[0])

		if !ok {
			return status.Error(
				codes.Unauthenticated,
				"unauthorized",
			)
		}

		headers.Append("number", number)

		ctx := metadata.NewIncomingContext(stream.Context(), headers)
		return handler(srv, &serverStream{stream, ctx})
	}
}
