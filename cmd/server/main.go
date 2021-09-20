package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	pb "github.com/maxkalayda/lnkCutNew/api/proto"
	"github.com/maxkalayda/lnkCutNew/pkg/handler"
	"github.com/maxkalayda/lnkCutNew/pkg/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"math/rand"
	"net"
	"time"
)

const (
	port = ":50051"
)

func main() {
	rand.Seed(time.Now().Unix())

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("env not loaded: %s", err.Error())
	}

	db, err := repository.PostgresConnect()
	if err != nil {
		log.Fatalf("failed to init db: %s, %s", err.Error(), db)
	}
	//test

	//test
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLinkServiceServer(s, &handler.Server{})
	//для тестов
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
