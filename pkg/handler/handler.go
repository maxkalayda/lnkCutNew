package handler

import (
	"context"
	"github.com/maxkalayda/lnkCutNew/api/proto"
	"github.com/maxkalayda/lnkCutNew/pkg/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"unicode/utf8"
)

//type Server struct {
//	proto.UnimplementedLinkServiceServer
//}

type Server struct {
	proto.UnimplementedLinkServiceServer
	Lr        *service.Implementation
	ShortLink string
	OrigLink  string
}

func (s *Server) Create(ctx context.Context, in *proto.LinkRequest) (*proto.LinkReply, error) {
	tmpOrig := in.GetLink()
	log.Printf("Server | Received from client origLink: %v", in.GetLink())
	//вызов записи и обработки линка
	tmpShort := service.CuttingLink(tmpOrig)
	log.Println("###Start work with DB")
	//log.Printf("Lr3 %v\n", s.Lr)
	err := s.Lr.LinkRepo.AddLink(tmpShort, tmpOrig)
	if err != nil {
		return &proto.LinkReply{Url: tmpShort}, err
	}
	log.Println("###End work with DB")
	return &proto.LinkReply{Url: tmpShort}, nil
}

func (s *Server) Get(ctx context.Context, in *proto.LinkRequest) (*proto.LinkReply, error) {
	tmp := in.GetLink() //получили укороченную линку
	tmpLen := utf8.RuneCountInString(tmp)
	log.Printf("START: Get from client short link: %s", tmp)
	//_, ok := pkg.MSync.Load(tmp)
	//SearchShort, SearchOrig, _ :=s.Lr.LinkRepo.SearchRow(tmp)
	//log.Printf("Get %v, %v", SearchShort, SearchOrig)
	var tmpOrig string

	if tmpLen != 10 {
		err := status.Newf(
			codes.InvalidArgument,
			"Длина ссылки не равна 10 символов.")
		log.Println("Длина ссылки не равна 10 символов")
		err, withDet := err.WithDetails(in)
		if withDet != nil {
			return nil, withDet
		}
		return nil, err.Err()
		//} else if err {
		//	err := status.Newf(
		//		codes.NotFound,
		//		"Укороченная ссылка не найдена")
		//	log.Println("Укороченная ссылка не найдена")
		//	err, withDet := err.WithDetails(in)
		//	if withDet != nil {
		//		return nil, withDet
		//	}
		//	return nil, err.Err()
	} else {
		//newVal, _ := pkg.MSync.Load(tmp)
		//tmp = newVal.(string)
		SearchShort, SearchOrig, err := s.Lr.LinkRepo.SearchRow(tmp)
		log.Println("err", err)
		log.Printf("END: Get from client short link: %s, %s", SearchShort, SearchOrig)
		tmpOrig = SearchOrig
		log.Println("orig len:", utf8.RuneCountInString(tmpOrig))
		log.Println("short len:", utf8.RuneCountInString(tmp))
		if err == nil && (utf8.RuneCountInString(tmpOrig) == 0 || utf8.RuneCountInString(tmp) == 0) {
			err := status.Newf(
				codes.NotFound,
				"Укороченная ссылка не найдена")
			log.Println("Укороченная ссылка не найдена")
			err, withDet := err.WithDetails(in)
			if withDet != nil {
				return nil, withDet
			}
			return nil, err.Err()
		}
	}
	return &proto.LinkReply{Url: tmpOrig}, nil
}
