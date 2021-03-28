package handler

import "github.com/gin-gonic/gin"

type document struct{}

var Document document


func (_ *document) GetByShortLink(c *gin.Context) {
	ErrorMsgResponse(c, "Not implemented")
}


func (_ *document) GetByQuery(c *gin.Context) {
	ErrorMsgResponse(c, "Not implemented")
}


func (_ *document) Create(c *gin.Context) {
	ErrorMsgResponse(c, "Not implemented")
}
