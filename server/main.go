package main

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	pb "grpc-todo/proto"
	"log"
	"net"
)

type TodoServer struct {
	pb.UnimplementedTodoServiceServer
}

func (s *TodoServer) CreateTodo(ctx context.Context, in *pb.NewTodo) (*pb.Todo, error) {
	log.Printf("Recived : %v", in.GetName())
	todo := &pb.Todo{
		Name: in.GetName(),
		Description: in.GetDescription(),
		Done: false,
		Id: uuid.New().String(),
	}

	return todo, nil;
}

func main() {
	const PORT = ":50051"

	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("Failed connection : ",err)
	}

	s := grpc.NewServer()

	pb.RegisterTodoServiceServer(s, &TodoServer{})

	log.Printf("Server listening at %v ",lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
