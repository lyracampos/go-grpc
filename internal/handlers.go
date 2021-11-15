package internal

import (
	"context"
	"fmt"

	pkg "github.com/lyracampos/go-grpc/pkg/pb/greeting/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GreeterService struct {
}

func (gs GreeterService) Greet(ctx context.Context, req *pkg.GreetRequest) (res *pkg.GreetResponse, err error) {
	if req.Message == nil {
		err = status.New(codes.InvalidArgument, "Message cannot be empty.").Err()
		return
	}

	helloMsg := fmt.Sprintf("%s, %s", req.Message.Greeting.String(), req.Message.Name)

	res = &pkg.GreetResponse{
		Response: helloMsg,
	}
	return
}
