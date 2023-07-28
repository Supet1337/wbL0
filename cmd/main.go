package main

import (
	"log"
	"wb-l0/config"
	"wb-l0/internal/handlers"
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
	server, err := handlers.CreateServer(conf)
	if err != nil {
		log.Fatal("Cant starting server", err)
	}
	err = server.StartServer(":3000")
	if err != nil {
		log.Fatalf("Cant starting server on this port", err)
	}
	log.Println("Starting server...")
}
