package main

import (
	//"context"
	"github.com/docker/docker/client"
	"github.com/simpleminepanel/simpleminepanel/routes"
	"log"
)

func main() {
	logger := log.Default()

	//dockerCtx := context.Background()
	dockerCli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic("could not connect to docker; please make sure it is running")
	}
	defer dockerCli.Close()
	router := routes.InitRouter()
	router.Static("/assets", "./assets")
	logger.Println("SMP is starting up...")
	router.Run()
}
