package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/utmhikari/protobuf-grpc-starter/internal/shared/models"
	service "github.com/utmhikari/protobuf-grpc-starter/internal/svr/websvr/service/document"
)

type document struct{}

var Document document


func (_ *document) GetByShortLink(c *gin.Context) {
	shortLink := c.Param("short-link")
	doc, err := service.GetDocumentByShortLink(shortLink)
	if err != nil {
		ErrorResponse(c, err)
		return
	}
	SuccessDataResponse(c, doc)
}


func (_ *document) GetByQuery(c *gin.Context) {
	var query models.Query
	if err := c.ShouldBindQuery(&query); err != nil {
		ErrorResponse(c, err)
		return
	}

	docs, err := service.GetDocumentsByQuery(&query)
	if err != nil {
		ErrorResponse(c, err)
	}

	SuccessDataResponse(c, docs)
}


func (_ *document) Create(c *gin.Context) {
	var doc models.Document
	if err := c.ShouldBindJSON(&doc); err != nil {
		ErrorResponse(c, err)
		return
	}

	err := service.CreateDocument(&doc)
	if err != nil {
		ErrorResponse(c, err)
		return
	}

	SuccessMsgResponse(c, "create doc successfully")
}
