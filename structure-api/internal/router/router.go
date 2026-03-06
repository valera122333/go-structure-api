package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"org-structure-api/internal/handler"
)

func NewRouter(depHandler *handler.DepartmentHandler) http.Handler {

	r := chi.NewRouter()

	r.Post("/departments", depHandler.CreateDepartment)
	r.Post("/departments/{id}/employees", depHandler.CreateEmployee)
	r.Get("/departments/{id}", depHandler.GetDepartment)

	return r
}
