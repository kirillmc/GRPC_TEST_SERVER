package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/kirillmc/data_filler/pkg/filler_pb"
	"github.com/kirillmc/grpc_test_server/internal/converter"
	desc "github.com/kirillmc/grpc_test_server/pkg/program_v3"
)

type server struct {
	desc.UnimplementedProgramV3Server
}

const messageSize = 1024 * 1024 * 1024

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
	//grpc.MaxRecvMsgSize(messageSize),
	)
	reflection.Register(s)
	desc.RegisterProgramV3Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.TrainPrograms, error) {
	// Дополнение:
	if req.GetCount() == 0 {
		return nil, errors.New("count is 0")
	}

	var programs = filler_pb.CreateOwnSetOfPrograms(int(req.GetCount()))

	return converter.ToResponseProgramsFromRepo(programs), nil
}
