package main

import (
	"log"

	"github.com/ARYPROGRAMMER/GoEmailApi/internal/env"
	"github.com/joho/godotenv"
)

func main(){
	if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	app := &application{
		config: cfg,
	}

	mux :=app.mount()
	log.Fatal(app.run(mux))
}