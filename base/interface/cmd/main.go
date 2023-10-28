package cmd

import (
	"github.com/Snaptime23/snaptime-server/v2/base/interface/internal/server/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const PORT = "9001"

func Run() {
	conn, err := grpc.Dial(":"+PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	if err != nil {
		panic(err)
	}

	http.NewServer(conn)
}
