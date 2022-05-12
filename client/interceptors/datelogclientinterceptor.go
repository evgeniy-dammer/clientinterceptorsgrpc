package interceptors

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
)

func DateLogClientInterceptor(ctx context.Context, method string, req interface{}, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	today := time.Now()

	fmt.Println("Client call services date: " + today.Format("01/02/2006 15:04:05"))

	err := invoker(ctx, method, req, reply, cc, opts...)

	return err
}
