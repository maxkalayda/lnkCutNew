package handler

import (
	"context"
	"github.com/maxkalayda/lnkCutNew/api/proto"
	"github.com/maxkalayda/lnkCutNew/pkg/service"
	lnkCutNew "github.com/maxkalayda/lnkCutterNew/pkg/service"
	"log"
	"unicode/utf8"
)

type Server struct {
	proto.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *Server) Create(ctx context.Context, in *proto.LinkRequest) (*proto.LinkReply, error) {
	tmp := in.GetName()
	log.Printf("Server | Received from client origLink: %v", in.GetName())
	tmp = lnkCutNew.CuttingLink(tmp)
	//tmp = CuttingLink(tmp)
	return &proto.LinkReply{Message: "Server | Client get short link: " + tmp}, nil
}

// SayHello implements helloworld.GreeterServer
func (s *Server) Get(ctx context.Context, in *proto.LinkRequest) (*proto.LinkReply, error) {
	tmp := in.GetName()
	_, ok := service.DbMap[tmp]
	if utf8.RuneCountInString(tmp) < 10 {
		log.Println("Длина ссылки меньше 10 символов")
		return &proto.LinkReply{Message: "Длина ссылки меньше 10"}, nil
	} else if !ok {
		log.Println("Укороченная ссылка не найдена", ok)
		return &proto.LinkReply{Message: "Укороченная ссылка не найдена"}, nil
	} else {
		log.Printf("Received from client short link: %v", in.GetName())
		tmp = service.DbMap[tmp]
	}
	return &proto.LinkReply{Message: "Server | Original Link: " + tmp}, nil
}
