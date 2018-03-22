package main

import (
	"os"
	"strconv"
	"time"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	pb "bitbucket.com/go"

)

const (
	address     = "localhost:50051"
)

func main(){
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewFizzbuzzClient(conn)

	// Contact the server and print out its response.
	var string1 string = os.Args[1]
	var string2 string = os.Args[2]

	int1, errInt1 := strconv.Atoi(os.Args[3])
	if errInt1 != nil {
		log.Fatal(errInt1)
	}
	int2, errInt2 := strconv.Atoi(os.Args[4])
	if errInt2 != nil {
		log.Fatal(errInt2)
	}
	limit, errLimit := strconv.Atoi(os.Args[5])
	if errLimit != nil {
		log.Fatal(errLimit)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.DoFizzbuzz(ctx, &pb.FizzbuzzRequest{
		String1: string1,
		String2: string2,
		Int1: int32(int1),
		Int2: int32(int2),
		Limit: int32(limit),
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	for _, content := range r.Reply {
		println(content)
	}
}
