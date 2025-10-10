package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	pb "github.com/jdk829355/learn_golang/go_grpc/helloworld_pb"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	"google.golang.org/genai"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type GongGamServer struct {
	pb.UnimplementedGongGamServer
}

func CreateMessage(s string) pb.GongGamResponse {
	return pb.GongGamResponse{Gonggam: s}
}

func (s *GongGamServer) YesYes(stream pb.GongGam_YesYesServer) error {
	apiKey := os.Getenv("GEMINI_API_KEY")
	end := CreateMessage("<end>")
	if apiKey == "" {
		log.Fatal("GEMINI_API_KEY environment variable not set")
	}

	// Create a new Gemini client.
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	for {
		m, err := stream.Recv()
		if err == io.EOF {
			return nil
		} else {
			if err != nil {
				return err
			}
			if m.GetGomin() == "bye" {
				return nil
			}
			log.Printf("received: %s", m.GetGomin())
			result := client.Models.GenerateContentStream(ctx, "gemini-2.5-flash", genai.Text("다음 문장에 공감하는 대답을 생성하여라. 최대한 따뜻한 말투로 감정적 공감을 할 것: "+m.GetGomin()), nil)

			for chunk, _ := range result {
				if len(chunk.Candidates) == 0 {
					break
				}
				part := chunk.Candidates[0].Content.Parts[0]
				messageToSend := CreateMessage(part.Text)
				err := stream.Send(&messageToSend)
				if err != nil {
					return err
				}
			}
			err := stream.Send(&end)
			if err != nil {
				return err
			}
		}
	}
}

func main() {
	flag.Parse()
	_ = godotenv.Load()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterGongGamServer(s, &GongGamServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
