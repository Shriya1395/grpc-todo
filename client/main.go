package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc-todo/proto"
	"log"
	"time"
)

type TodoTask struct {
	Name 		string
	Description string
	Done 		bool
}

func main() {
	const ADDR = "localhost:50051"

	conn, err := grpc.Dial(ADDR, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect : %v", err)
	}
	defer conn.Close()

	c := pb.NewTodoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	todos := []TodoTask{
		{Name: "Code review", Description: "Review new feature code", Done: false},
		{Name: "Make YouTube Video", Description: "Start Go for beginners series", Done: false},
		{Name: "Go to the gym", Description: "Leg day", Done: false},
		{Name: "Buy groceries", Description: "Buy tomatoes, onions, mangos", Done: false},
		{Name: "Meet with mentor", Description: "Discuss blockers in my project", Done: false},
	}

	for _, todo := range todos {
		res, err := c.CreateTodo(ctx, &pb.NewTodo{Name: todo.Name, Description: todo.Description, Done: todo.Done})
		if err != nil {
			log.Fatalf("could not create user: %v", err)
		}

		log.Printf(`
			ID : %s
			Name : %s 
			Desription : %s 
			Done : %v,
			`, res.GetId(), res.GetName(), res.GetDescription(), res.GetDone())
	}
}