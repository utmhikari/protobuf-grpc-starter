package config

import (
	"errors"
	"fmt"
	"github.com/utmhikari/protobuf-grpc-starter/internal/shared/util/fs"
	"log"
)

type ServerConfig struct {
	InternalPort int `json:"internalPort"`
	ExternalPort int `json:"externalPort"`
}

func (sc *ServerConfig) GetInternalPortStr() string {
	return fmt.Sprintf(":%d", sc.InternalPort)
}

func (sc *ServerConfig) GetExternalPortStr() string {
	return fmt.Sprintf(":%d", sc.ExternalPort)
}


type ClusterConfig struct {
	ServerConfigs map[string]ServerConfig `json:"serverConfigs"`
}


const CfgRoot string = "./configs"


func GetConfigPath(name string) string {
	return fmt.Sprintf("%s/%s", CfgRoot, name)
}


func GetClusterConfig() (*ClusterConfig, error) {
	cfgPath := GetConfigPath("cluster.json")
	cfg := ClusterConfig{}
	err := fs.ReadJsonFile(cfgPath, &cfg)
	if err != nil {
		return nil, err
	}

	// log.Printf("cluster config: %+v\n", cfg)
	return &cfg, nil
}


func GetServerConfig(name string) (*ServerConfig, error) {
	clusterCfg, err := GetClusterConfig()
	if err != nil {
		return nil, err
	}

	svrCfg, ok := clusterCfg.ServerConfigs[name]
	if !ok {
		return nil, errors.New("cannot find cfg of server " + name)
	}

	log.Printf("server %s config: %+v\n", name, svrCfg)

	return &svrCfg, nil
}


