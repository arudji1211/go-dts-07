package main

import "github.com/aruji1211/myGram/server"

// @host      localhost:9999
// @BasePath /api/v1
// @securityDefinitions.apiKey Bearer
// @in header
// @name Authorization

func main() {
	server.Serve()
}
