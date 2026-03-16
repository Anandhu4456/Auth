package main

import (
	"auth/config"
	"auth/db"
	"auth/router"

	"log"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Printf("load config error : %v", err)
	}

	pool, err := db.ConnectDB(cfg)
	if err != nil {
		log.Printf("db connection failed : %v", err)
	}

	defer pool.Close()

	log.Println("[DB] connection established..")

	engine := router.RouteController(pool, cfg)

	if err := engine.Run(":8080"); err != nil {
		log.Fatalf("server failed : %v", err)
	}

	log.Println("[Auth Service is running successfully...]")

}
