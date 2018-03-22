package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"strconv"
	"fmt"
	pb "bitbucket.org/zuful/fizz-buzz-grpc-protobuf-go"
)

const (
	port = ":50051"
)

type server struct{}

// DoFizzbuzz implements fizzbuzz.FizzbuzzServer
func (s *server) DoFizzbuzz(ctx context.Context, in *pb.FizzbuzzRequest) (*pb.FizzbuzzReply, error) {

	var list []string = make([]string, 0)

	for i := int32(1); i <= in.Limit; i++{

		if (i % in.Int1 == 0) && (i % in.Int2 == 0){

			list = append(list, in.String1 + in.String2)

		} else if i % in.Int1 == 0 {

			list = append(list, in.String1)

		} else if i % in.Int2 == 0 {

			list = append(list, in.String2)

		} else {

			list = append(list, strconv.Itoa(int(i)))

		}
	}

	return &pb.FizzbuzzReply{Reply:list}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterFizzbuzzServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	} else {
		fmt.Println("Serving on port: " + port)
	}
}
