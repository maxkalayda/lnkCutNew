/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
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
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &handler.Server{})
	//для тестов
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
