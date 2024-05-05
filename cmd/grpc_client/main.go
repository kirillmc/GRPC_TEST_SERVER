package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"

	desc "github.com/kirillmc/grpc_test_server/pkg/program_v3"
)

const (
	address     = "localhost:50051"
	oldAddress  = "2.tcp.eu.ngrok.io:10883"
	messageSize = 1024 * 1024 * 1024
)

// func DialogOptions(chains ...grpc.UnaryClientInterceptor) []grpc.DialOption {
func DialogOptions() []grpc.DialOption {
	//	chains = append(chains, userinfo.UnaryClientInterceptor())
	return []grpc.DialOption{
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(messageSize),
			grpc.MaxCallSendMsgSize(messageSize),
		),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		//grpc.WithChainUnaryInterceptor(chains...),
	}
}

func getNProgramsClient(n int64) (*desc.TrainPrograms, error) {
	conn, err := grpc.Dial(address,
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(messageSize)),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect %v", err)
		return nil, err
	}

	defer conn.Close()

	c := desc.NewProgramV3Client(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	//maxSizeOption := grpc.MaxCallRecvMsgSize(1024 * 1024 *1024)
	programs, err := c.Get(ctx, &desc.GetRequest{Count: n})
	if err != nil {
		log.Println(err)
	}

	return programs, nil
}

func main() {
	start := time.Now()
	programs, err := getNProgramsClient(55)
	if err != nil {
		log.Println(err)
	}
	end := time.Now()
	numOfSets, err := proto.Marshal(programs)
	if err != nil {
		fmt.Errorf("fail to get json: %v", err)
	}
	log.Printf("TOTAL TIME TO GET PROGRAMS: %v\n", end.Sub(start))
	log.Printf("SEZE OF PROGRAMS: %d\n", len(numOfSets))
}
