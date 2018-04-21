package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/jstrachan/mymicro4/handler"
	"github.com/jstrachan/mymicro4/subscriber"

	example "github.com/jstrachan/mymicro4/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.mymicro4"),
		micro.Version("latest"),
	)

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.mymicro4", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.mymicro4", service.Server(), subscriber.Handler)

	// Initialise service
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
