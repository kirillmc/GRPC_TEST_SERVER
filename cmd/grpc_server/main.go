package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	filler "github.com/kirillmc/grpc_test_server/internal/filler_pb"
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
		grpc.MaxSendMsgSize(messageSize),
	)
	reflection.Register(s)
	desc.RegisterProgramV3Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.TrainPrograms, error) {

	var programs = filler.CreateOwnSetOfPrograms(int(req.GetCount()))

	return programs, nil
}

func (s *server) Create(ctx context.Context, req *desc.TrainPrograms) (*desc.Response, error) {
	req.GetTrainPrograms()

	return &desc.Response{Message: "Данные были добавлены"}, nil
}

func (s *server) Update(ctx context.Context, req *desc.TrainPrograms) (*desc.Response, error) {
	req.GetTrainPrograms()

	return &desc.Response{Message: "Данные были обновлены"}, nil
}

func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*desc.Response, error) {
	req.GetId()

	return &desc.Response{Message: "Данные были удалены"}, nil
}
