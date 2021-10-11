package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	pb "github.com/maxkalayda/lnkCutNew/api/proto"
	"github.com/maxkalayda/lnkCutNew/pkg/handler"
	"github.com/maxkalayda/lnkCutNew/pkg/repository"
	"github.com/maxkalayda/lnkCutNew/pkg/service"
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

	Lr := service.NewLinkRepo(db)
	s := grpc.NewServer()
	srv := &handler.Server{Lr: Lr}
	pb.RegisterLinkServiceServer(s, srv)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
