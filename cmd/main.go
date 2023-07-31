package main

import (
	"log"
	"wb-l0/config"
	"wb-l0/internal/handlers"
	"wb-l0/internal/handlers/nats"
	"wb-l0/internal/usecase"
)

func main() {
	viperConf, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	conf, err := config.ParseConfig(viperConf)
	if err != nil {
		log.Fatal(err)
	}

	ucase, err := usecase.NewUsecase(conf)
	stan, err := nats.NewNats(ucase)
	if err != nil {
		log.Fatalf("can not connect to nats streaming: %v", err)
	}
	err = stan.Subscribe("topic")
	if err != nil {
		log.Fatalf("Cant subscribe on topic: %s. %v", "topic", err)
	}

	server, err := handlers.CreateServer(ucase)
	if err != nil {
		log.Fatal("Cant starting server", err)
	}
	err = server.StartServer(":3000")
	if err != nil {
		log.Fatalf("Cant starting server on this port", err)
	}
	log.Println("Starting server...")
}
