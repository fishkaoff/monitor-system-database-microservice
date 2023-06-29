package main

import (
	"context"
	"fmt"
	"log"
	"os"

	grpcserver "github.com/fishkaoff/monitor-system-database-microservice/grpcServer"
	"github.com/fishkaoff/monitor-system-database-microservice/service"
	storage "github.com/fishkaoff/monitor-system-database-microservice/storage/postgresql"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)


func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Cannot load env file")
	}
}


func main() {
	DBUrl := os.Getenv("DB_URL")

	conn, err := pgxpool.New(context.Background(), DBUrl)
	if err != nil {
		log.Fatal(err)
	}

	storage := storage.NewDB(conn)
	serv := service.New(storage)


	listenAddr := ":4000"
	fmt.Println("Server is running")
	log.Fatal(grpcserver.GRPCServerRun(listenAddr, serv))
}