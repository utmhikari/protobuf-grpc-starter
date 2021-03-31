package cache

import (
	"context"
	"errors"
	pb "github.com/utmhikari/protobuf-grpc-starter/api/pb/base"
	pbCache "github.com/utmhikari/protobuf-grpc-starter/api/pb/cache"
	"github.com/utmhikari/protobuf-grpc-starter/internal/shared/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	"net"
	"time"
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
	document := Get(in.ShortLink)
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

	Set(in.Document)
	return &pbCache.SetDocumentResponse{
		Status: &pb.RespStatus{
			Success: true,
			Message: "",
		},
	}, nil
}


// grpc-go/examples/features/keepalive/server/main.go
var kaep = keepalive.EnforcementPolicy{
	MinTime:             5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
	PermitWithoutStream: true,            // Allow pings even when there are no active streams
}

// grpc-go/examples/features/keepalive/server/main.go
var kasp = keepalive.ServerParameters{
	MaxConnectionIdle:     15 * time.Second, // If a client is idle for 15 seconds, send a GOAWAY
	MaxConnectionAge:      30 * time.Second, // If any connection is alive for more than 30 seconds, send a GOAWAY
	MaxConnectionAgeGrace: 5 * time.Second,  // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
	Time:                  5 * time.Second,  // Ping the client if it is idle for 5 seconds to ensure the connection is still active
	Timeout:               1 * time.Second,  // Wait 1 second for the ping ack before assuming the connection is dead
}


// startServer see helloworld example of grpc-go
func StartServer(cfg *config.ServerConfig) {
	lis, err := net.Listen("tcp", cfg.GetInternalPortStr())
	if err != nil {
		log.Fatalf("cachesvr: failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.KeepaliveEnforcementPolicy(kaep),
		grpc.KeepaliveParams(kasp))
	pbCache.RegisterCacheServer(s, &server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("cachesvr: failed to serve: %v", err)
	}
}