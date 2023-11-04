package cmd

import (
	"github.com/Snaptime23/snaptime-server/v2/mw"
	"github.com/Snaptime23/snaptime-server/v2/video/interface/internal/server/http"
	"github.com/Snaptime23/snaptime-server/v2/video/interface/internal/server/router"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

const PORT = "9002"

func Run() {
	conn, err := grpc.Dial(":"+PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	if err != nil {
		panic(err)
	}

	server := http.NewServer(conn)

	app := router.CreateEngine()
	app.Use(mw.Cors())
	api := app.Group("/api")
	router.InitVideoRouter(api, server)
	app.SetTrustedProxies(nil)

	// first use env, then default
	// app must listen on 9000 while running on tencent scf
	host := os.Getenv("HOST")
	if host == "" {
		host = "127.0.0.1"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "3006"
	}
	if os.Getenv("SERVERLESS") != "" {
		host = "0.0.0.0"
		port = "10000"
	}

	println("[Snaptime] listening on " + host + ":" + port)
	app.Run(host + ":" + port)
}
