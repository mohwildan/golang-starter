package handler

import "github.com/gin-gonic/gin"

type SampleHandler interface {
	RegisterRoutes(router *gin.Engine)
	List(c *gin.Context)
	Create(c *gin.Context)
	Detail(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
