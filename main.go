package main

import "github.com/arudji1211/go-dts-07/server"

// @host      localhost:9999
// @BasePath /api/v1
// @securityDefinitions.apiKey Bearer
// @in header
// @name Authorization

func main() {
	server.Serve()
}
