package main

import (
	"os"

	"github.com/Snaptime23/snaptime-server/v2/router"
)

func main() {
	app := router.CreateEngine()
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
