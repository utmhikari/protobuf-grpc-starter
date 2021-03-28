package cachesvr

import (
	"context"
	"errors"
	pb "github.com/utmhikari/protobuf-grpc-starter/api/pb/base"
	pbCache "github.com/utmhikari/protobuf-grpc-starter/api/pb/cache"
	"github.com/utmhikari/protobuf-grpc-starter/internal/shared/config"
	"github.com/utmhikari/protobuf-grpc-starter/internal/svr/cachesvr/cache"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pbCache.UnimplementedCacheServer
}


func (s *server) GetDocument(ctx context.Context, in *pbCache.GetDocumentRequest) (*pbCache.GetDocumentResponse, error) {
	if nil == in {
		return nil, errors.New("nil request")
	}

	if "" == in.ShortLink {
		return nil, errors.New("empty request")
	}

	log.Printf("get document: %+v\n", in)
	document := cache.Get(in.ShortLink)
	if nil == document {
		return &pbCache.GetDocumentResponse{
			Status: &pb.RespStatus{
				Success: false,
				Message: "cannot find document",
			},
			Document: nil,
		}, nil
	}

	return &pbCache.GetDocumentResponse{
		Status: &pb.RespStatus{
			Success: true,
			Message: "",
		},
		Document: document,
	}, nil
}


func (s *server) SetDocument(ctx context.Context, in *pbCache.SetDocumentRequest) (*pbCache.SetDocumentResponse, error) {
	if nil == in || nil == in.Document {
		return nil, errors.New("nil request")
	}

	cache.Set(in.Document)
	return &pbCache.SetDocumentResponse{
		Status: &pb.RespStatus{
			Success: true,
			Message: "",
		},
	}, nil
}


// startServer see helloworld example of grpc-go
func startServer(cfg *config.ServerConfig) {
	lis, err := net.Listen("tcp", cfg.GetInternalPortStr())
	if err != nil {
		log.Fatalf("cachesvr: failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pbCache.RegisterCacheServer(s, &server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("cachesvr: failed to serve: %v", err)
	}
}