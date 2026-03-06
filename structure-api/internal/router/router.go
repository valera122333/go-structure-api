package router

import (
	"net/http"

	"org-structure-api/internal/handler"

	"github.com/go-chi/chi/v5"
)

func NewRouter(depHandler *handler.DepartmentHandler) http.Handler {
	r := chi.NewRouter()

	// CRUD департаментов
	r.Post("/departments", depHandler.CreateDepartment)
	r.Post("/departments/{id}/employees", depHandler.CreateEmployee)
	r.Get("/departments/{id}", depHandler.GetDepartment)
	r.Patch("/departments/{id}", depHandler.UpdateDepartment)
	r.Delete("/departments/{id}", depHandler.DeleteDepartment)

	return r
}
