package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/devfullcycle/gRPC/internal/database"
	"github.com/devfullcycle/gRPC/internal/pb"
	"github.com/devfullcycle/gRPC/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	categoryDb := database.NewCategory(db)
	categoryServer := service.NewCategoryService(*categoryDb) //serviço de categoria

	//Criação do servidor gRPC
	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryServer) //registro services no servidor gRPC
	//ler e processa a sua propria informação
	reflection.Register(grpcServer)

	//Abre conexão tcp na porta 50051
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
