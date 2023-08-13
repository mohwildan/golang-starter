package handler

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type SampleHandler interface {
	RegisterRoutes(router *chi.Mux)
	List(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Detail(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}
