package cmd

import (
	"github.com/Snaptime23/snaptime-server/v2/base/interface/internal/server/http"
	"github.com/Snaptime23/snaptime-server/v2/base/interface/internal/server/router"
	"github.com/Snaptime23/snaptime-server/v2/mw"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

const BasePORT = "9001"
const VideoPORT = "9002"

func Run() {
	connBase, err := grpc.Dial(":"+BasePORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	connVideo, err := grpc.Dial(":"+VideoPORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer func() {
		connBase.Close()
		connVideo.Close()
	}()

	server := http.NewServer(connBase, connVideo)

	app := router.CreateEngine()
	app.Use(mw.Cors())
	api := app.Group("/api")
	router.InitBaseRouter(api, server)
	app.SetTrustedProxies(nil)

	// first use env, then default
	// app must listen on 9000 while running on tencent scf
	host := os.Getenv("HOST")
	if host == "" {
		host = "127.0.0.1"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "3005"
	}
	if os.Getenv("SERVERLESS") != "" {
		host = "0.0.0.0"
		port = "9000"
	}

	println("[Snaptime] listening on " + host + ":" + port)
	app.Run(host + ":" + port)
}
