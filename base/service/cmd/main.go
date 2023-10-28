package main

import (
	"github.com/Snaptime23/snaptime-server/v2/base/internal/api"
	"github.com/Snaptime23/snaptime-server/v2/base/service/internal/server"
	"google.golang.org/grpc"
	"net"
)

const PORT = "9001"

func main() {
	s := grpc.NewServer()
	api.RegisterBaseServiceServer(s, &server.Server{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		panic(err)
	}
	s.Serve(lis)
}
