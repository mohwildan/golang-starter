package handler

import (
	"golang-starter/service/domain/entity"
	"golang-starter/service/domain/usecase"
	"log"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

type sampleHandler struct {
	usecase usecase.SampleUsecase
}

func SampleHandler(g *gin.Engine, uc usecase.SampleUsecase) {
	handler := sampleHandler{
		usecase: uc,
	}
	v1 := g.Group("/sample/v1")
	v1.GET("/list", handler.List)
	v1.GET("/detail/:id", handler.Detail)
	v1.PUT("/update/:id", handler.Update)
	v1.POST("/create", handler.Create)
	v1.DELETE("/delete/:id", handler.Delete)
}

func (h *sampleHandler) List(c *gin.Context) {
	ctx := c.Request.Context()
	options := map[string]interface{}{
		"query": c.Request.URL.Query(),
	}
	response := h.usecase.List(ctx, options)
	c.JSON(response.Status, response)
}

func (h *sampleHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var request entity.SampleMongo
	err := json.NewDecoder(c.Request.Body).Decode(&request)
	if err != nil {
		log.Println(err.Error())
	}
	options := map[string]interface{}{
		"request": request,
	}
	reponse := h.usecase.Create(ctx, options)
	c.JSON(reponse.Status, reponse)
}

func (h *sampleHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	options := map[string]interface{}{
		"id": c.Param("id"),
	}
	reponse := h.usecase.Delete(ctx, options)
	c.JSON(reponse.Status, reponse)
}

func (h *sampleHandler) Detail(c *gin.Context) {
	ctx := c.Request.Context()
	options := map[string]interface{}{
		"id": c.Param("id"),
	}
	reponse := h.usecase.Detail(ctx, options)
	c.JSON(reponse.Status, reponse)
}

func (h *sampleHandler) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var request entity.SampleMongo
	err := json.NewDecoder(c.Request.Body).Decode(&request)
	if err != nil {
		log.Println(err.Error())
	}
	options := map[string]interface{}{
		"id":      c.Param("id"),
		"request": request,
	}
	reponse := h.usecase.Update(ctx, options)
	c.JSON(reponse.Status, reponse)
}
