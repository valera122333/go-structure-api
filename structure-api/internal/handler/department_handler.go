package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"org-structure-api/internal/service"
)

type DepartmentHandler struct {
	service *service.DepartmentService
}

func NewDepartmentHandler(s *service.DepartmentService) *DepartmentHandler {
	return &DepartmentHandler{s}
}

func (h *DepartmentHandler) CreateDepartment(w http.ResponseWriter, r *http.Request) {

	var req struct {
		Name     string `json:"name"`
		ParentID *uint  `json:"parent_id"`
	}

	json.NewDecoder(r.Body).Decode(&req)

	dep, err := h.service.CreateDepartment(req.Name, req.ParentID)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(dep)
}

func (h *DepartmentHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)

	var req struct {
		FullName string `json:"full_name"`
		Position string `json:"position"`
	}

	json.NewDecoder(r.Body).Decode(&req)

	emp, err := h.service.CreateEmployee(uint(id), req.FullName, req.Position)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(emp)
}

func (h *DepartmentHandler) GetDepartment(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)

	dep, err := h.service.GetDepartmentTree(uint(id), 3)

	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	json.NewEncoder(w).Encode(dep)
}
