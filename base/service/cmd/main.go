package cmd

import (
	"github.com/Snaptime23/snaptime-server/v2/base/rpc_pb/baseApi"
	"github.com/Snaptime23/snaptime-server/v2/base/service/internal/dao"
	"github.com/Snaptime23/snaptime-server/v2/base/service/internal/server"
	"google.golang.org/grpc"
	"net"
)

const PORT = "9001"

func Run() {
	dao.Init()

	s := grpc.NewServer()
	baseApi.RegisterBaseServiceServer(s, server.NewServer())

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		panic(err)
	}
	s.Serve(lis)
}
