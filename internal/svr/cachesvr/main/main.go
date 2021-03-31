package main

import (
	"github.com/utmhikari/protobuf-grpc-starter/internal/shared/config"
	"github.com/utmhikari/protobuf-grpc-starter/internal/shared/util/fs"
	"github.com/utmhikari/protobuf-grpc-starter/internal/svr/cachesvr/cache"
	"log"
)


func main() {
	cfgPath := config.GetConfigPath("cachesvr.json")
	cfg := cache.Config{}
	err := fs.ReadJsonFile(cfgPath, &cfg)
	if err != nil {
		panic(err)
	}

	svrCfg, err := config.GetServerConfig("cachesvr")
	if err != nil {
		panic(err)
	}

	log.Printf("launch cachesvr with svrCfg: %+v; cacheCfg: %+v\n", svrCfg, cfg)

	err = cache.Init(&cfg)
	if err != nil {
		panic(err)
	}

	cache.StartServer(svrCfg)
}