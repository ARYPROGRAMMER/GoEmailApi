package main

import (
	"log"

	"github.com/ARYPROGRAMMER/GoEmailApi/internal/db"
	"github.com/ARYPROGRAMMER/GoEmailApi/internal/env"
	"github.com/ARYPROGRAMMER/GoEmailApi/internal/store"
	"github.com/joho/godotenv"
)

func main(){
	if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }



	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr: env.GetString("DB_ADDR", "postgres://postgres:password@localhost/email?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime: env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}
	
	db,err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Print("Database connection pool established")

	store := store.PostgresStorage(db)

	app := &application{
		config: cfg,
		store: store,
	}


	mux :=app.mount()
	log.Fatal(app.run(mux))
}

