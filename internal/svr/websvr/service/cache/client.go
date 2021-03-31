package cache

import (
	"fmt"
	"github.com/utmhikari/protobuf-grpc-starter/api/pb/cache"
	"github.com/utmhikari/protobuf-grpc-starter/internal/shared/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	"time"
)

// websvr -> cachesvr client

var cacheSvrClient cache.CacheClient

// grpc-go/examples/features/keepalive/client/main.go
var kacp = keepalive.ClientParameters{
	Time:                10 * time.Second,
	Timeout:             time.Second,
	PermitWithoutStream: true,
}

func StartCacheSvrClient() error {
	cacheSvrCfg, err := config.GetServerConfig("cachesvr")
	if err != nil {
		return err
	}

	addr := fmt.Sprintf("localhost:%d", cacheSvrCfg.InternalPort)

	log.Printf("cache client -> connecting to %s\n", addr)

	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithKeepaliveParams(kacp))
	if err != nil {
		return err
	}

	go func() {
		cacheSvrClient = cache.NewCacheClient(conn)
		log.Printf("cache client connected~")
		// TODO: conn close err handling
	}()

	return nil
}

func GetClient() cache.CacheClient {
	return cacheSvrClient
}
