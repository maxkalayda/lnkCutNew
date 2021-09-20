package handler

import (
	"context"
	"github.com/maxkalayda/lnkCutNew/api/proto"
	"github.com/maxkalayda/lnkCutNew/pkg/service"
	lnkCutNew "github.com/maxkalayda/lnkCutNew/pkg/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"unicode/utf8"
)

type Server struct {
	proto.UnimplementedLinkServiceServer
}

func (s *Server) Create(ctx context.Context, in *proto.LinkRequest) (*proto.LinkReply, error) {
	tmp := in.GetLink()
	log.Printf("Server | Received from client origLink: %v", in.GetLink())
	tmp = lnkCutNew.CuttingLink(tmp)
	return &proto.LinkReply{Url: tmp}, nil
	//return &proto.LinkReply{Name2: in.GetName() + ":" + tmp}, nil
}

func (s *Server) Get(ctx context.Context, in *proto.LinkRequest) (*proto.LinkReply, error) {
	tmp := in.GetLink()
	tmpLen := utf8.RuneCountInString(tmp)
	_, ok := service.DbMap[tmp]
	if tmpLen < 10 {
		err := status.Newf(
			codes.InvalidArgument,
			"Длина ссылки меньше 10 символов.")
		log.Println("Длина ссылки меньше 10 символов")
		err, withDet := err.WithDetails(in)
		if withDet != nil {
			return nil, withDet
		}
		return nil, err.Err()
	} else if !ok {
		err := status.Newf(
			codes.NotFound,
			"Укороченная ссылка не найдена")
		log.Println("Укороченная ссылка не найдена")
		err, withDet := err.WithDetails(in)
		if withDet != nil {
			return nil, withDet
		}
		return nil, err.Err()
	} else {
		log.Printf("Получили от клиента укороченную ссылку: %v", in.GetLink())
		tmp = service.DbMap[tmp]
	}
	return &proto.LinkReply{Url: tmp}, nil
	//return &proto.LinkReply{Message: "Server | Original Link: " + tmp}, nil
}
