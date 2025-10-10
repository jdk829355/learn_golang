package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	pb "github.com/jdk829355/learn_golang/go_grpc/helloworld_pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("%v", err)
		}
	}(conn)

	gonggam := pb.NewGongGamClient(conn)

	// Contact the server and print out its response.
	ctx := context.Background()

	stream, err := gonggam.YesYes(ctx)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	w := make(chan struct{})

	go func() {
		for {
			reader := bufio.NewReader(os.Stdin)
			var q string

			fmt.Print("고민을 털어놓으세요: ")
			q, err := reader.ReadString('\n')
			q = strings.TrimSpace(q)
			if err != nil {
				log.Fatalf("%v", err)
			}
			err = stream.Send(&pb.GongGamRequest{Gomin: q})
			if err != nil {
				log.Fatalf("%v", err)
			}
			<-w
		}
	}()

	for {
		message, err := stream.Recv()
		if err != nil {
			log.Fatalf("%v", err)
		}
		if message.GetGonggam() == "<end>" {
			fmt.Print("\n")
			w <- struct{}{}
		} else {
			fmt.Print(message.GetGonggam())
		}
	}
}
