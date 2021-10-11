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
	//db connect
	db, err := repository.PostgresConnect()
	if err != nil {
		log.Fatalf("failed to init db: %s, %s", err.Error(), db)
	}

	//
	//Lr:=service.NewLinkRepo(db)
	//log.Printf("Lr %T\n", Lr)
	//Lr.LinkRepo.AddLink("444","224442")

	Lr := service.NewLinkRepo(db)
	Lr2 := service.NewLinkRepo(db)
	log.Printf("Lr %T\n", Lr)
	log.Printf("Lr2 %T\n", Lr2)

	//Lr.LinkRepo2.AddLink("555","555")
	//Mr := service.NewLinkRepoMap(pkg.MSync)

	//

	s := grpc.NewServer()
	srv := &handler.Server{}
	pb.RegisterLinkServiceServer(s, srv)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//для тестов
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
