package main

import (
	"context"
	"fmt"
	"go.uber.org/dig"
	"go_id/services/api/rest"
	"go_id/services/repo/product_service"
	"go_id/services/usecases"
	"os"
)

func main() {
	finished := make(chan bool)
	container := BuildContainer()
	err := container.Invoke(func(server rest.HttpServer) {
		server.Run(context.Background())
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Main: Waiting for worker to finish")
	<-finished
	fmt.Println("Main: Completed")
	os.Exit(0)
}

func BuildContainer() *dig.Container {
	container := dig.New()
	container.Provide(product_service.NewProduct)
	container.Provide(usecases.NewService)
	container.Provide(rest.NewHttpServer)
	return container
}
