package interceptors

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func MethodLogClientInterceptor(ctx context.Context, method string, req interface{}, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	fmt.Println("Client call " + method + "service")

	err := invoker(ctx, method, req, reply, cc, opts...)

	return err
}
