package main

import (
	"log"

	"github.com/ARYPROGRAMMER/GoEmailApi/internal/env"
	"github.com/ARYPROGRAMMER/GoEmailApi/internal/store"
	"github.com/joho/godotenv"
)

func main(){
	if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }

	store := store.PostgresStorage(nil)

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	app := &application{
		config: cfg,
		store: store,
	}


	mux :=app.mount()
	log.Fatal(app.run(mux))
}