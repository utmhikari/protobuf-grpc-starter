package websvr

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/utmhikari/protobuf-grpc-starter/internal/shared/config"
	"github.com/utmhikari/protobuf-grpc-starter/internal/svr/websvr/handler"
	"log"
	"net/http"
)


func getWebEngine() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		v1.GET("/health", handler.HealthCheck)

		documents := v1.Group("/documents")
		{
			documents.GET("/", handler.Document.GetByQuery)
			documents.POST("/", handler.Document.Create)
		}

		document := v1.Group("/document")
		{
			document.GET("/:short-link", handler.Document.GetByShortLink)
		}
	}

	return r
}


func Start() error {
	svrCfg, err := config.GetServerConfig("websvr")
	if err != nil {
		return err
	}

	engine := getWebEngine()
	if engine == nil {
		return errors.New("web engine is nil")
	}

	log.Printf("launch websvr with svrCfg: %+v\n", svrCfg)

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", svrCfg.ExternalPort),
		Handler: engine,
	}

	return httpServer.ListenAndServe()
}


func main() {
	err := Start()
	if err != nil {
		panic(err)
	}
}
