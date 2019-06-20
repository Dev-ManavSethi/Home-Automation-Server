package main

import (
	"context"
	"fmt"
	"net"

	"github.com/Dev-ManavSethi/Home-Automation-Server/services"

	"github.com/Dev-ManavSethi/Home-Automation-Server/utils"
	"google.golang.org/grpc"

	"github.com/joho/godotenv"
)

type Student struct {
	Name  string
	ID    int32
	Class string
}

type server struct{}

func init() {

	err := godotenv.Load(".env")
	utils.LogErrorOrSuccess(err, "Error setting env variables from /.env", "Sucessfully set env variables from /.env")

}

func (*server) CreateClient(ctx context.Context, CreateClientRequest *services.CreateClientRequest) (*services.CreateClientResponse, error) {

	return nil, nil
}

func main() {

	fmt.Println("Home Automation Server")

	Listener, err := net.Listen("tcp", "0.0.0.0:50051")
	utils.LogErrorOrSuccess(err, "Error listening at 0.0.0.0:50051", "gRPC server listening at 0.0.0.0:50051")

	Server := grpc.NewServer()
	services.RegisterControlHomeServer(Server, &server{})

	// Register reflection service on gRPC server.

	err02 := Server.Serve(Listener)
	utils.LogErrorOrSuccess(err02, "Error to serve the listener", "Successfully serving the Lister at 0.0.0.0:50051")

}
