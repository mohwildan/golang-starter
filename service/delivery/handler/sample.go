package handler

import (
	"context"
	"github.com/go-chi/chi/v5"
	"golang-starter/helpers/response"
	"golang-starter/service/domain/handler"
	"net/http"

	"golang-starter/service/domain/usecase"
)

type SampleHandler struct {
	usecase usecase.SampleUsecase
}

func NewSampleHandler(uc usecase.SampleUsecase) handler.SampleHandler {
	return &SampleHandler{
		usecase: uc,
	}
}

func (h *SampleHandler) RegisterRoutes(r *chi.Mux) {
	r.Route("/sample", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Get("/list", h.List)
		})
	})
}

func (h *SampleHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	options := map[string]any{
		"query": r.URL.Query(),
	}

	res := h.usecase.List(ctx, options)
	response.Json(res.Status, res, w)
}

func (h *SampleHandler) Create(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h *SampleHandler) Detail(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h *SampleHandler) Update(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h *SampleHandler) Delete(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
