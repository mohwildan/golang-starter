package handler

import (
	"golang-starter/service/domain/handler"
	"log"
	"net/http"

	"golang-starter/service/domain/entity"
	"golang-starter/service/domain/usecase"

	"github.com/gin-gonic/gin"
)

type SampleHandler struct {
	usecase usecase.SampleUsecase
}

func NewSampleHandler(uc usecase.SampleUsecase) handler.SampleHandler {
	return &SampleHandler{
		usecase: uc,
	}
}

func (h *SampleHandler) RegisterRoutes(router *gin.Engine) {
	v1 := router.Group("/sample/v1")
	v1.GET("/list", h.List)
	v1.GET("/detail/:id", h.Detail)
	v1.PUT("/update/:id", h.Update)
	v1.POST("/create", h.Create)
	v1.DELETE("/delete/:id", h.Delete)
}

func (h *SampleHandler) List(c *gin.Context) {
	ctx := c.Request.Context()
	options := map[string]any{
		"query": c.Request.URL.Query(),
	}
	response := h.usecase.List(ctx, options)
	c.JSON(response.Status, response)
}

func (h *SampleHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var request entity.SampleMongo
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	options := map[string]any{
		"request": request,
	}
	response := h.usecase.Create(ctx, options)
	c.JSON(response.Status, response)
}

func (h *SampleHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	options := map[string]any{
		"id": c.Param("id"),
	}
	response := h.usecase.Delete(ctx, options)
	c.JSON(response.Status, response)
}

func (h *SampleHandler) Detail(c *gin.Context) {
	ctx := c.Request.Context()
	options := map[string]any{
		"id": c.Param("id"),
	}
	response := h.usecase.Detail(ctx, options)
	c.JSON(response.Status, response)
}

func (h *SampleHandler) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var request entity.SampleMongo
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	options := map[string]any{
		"id":      c.Param("id"),
		"request": request,
	}
	response := h.usecase.Update(ctx, options)
	c.JSON(response.Status, response)
}
