package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"org-structure-api/internal/service"

	"github.com/go-chi/chi/v5"
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
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", 400)
		return
	}

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
		FullName string  `json:"full_name"`
		Position string  `json:"position"`
		HiredAt  *string `json:"hired_at"` // optional
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", 400)
		return
	}

	emp, err := h.service.CreateEmployee(uint(id), req.FullName, req.Position, req.HiredAt)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(emp)
}

func (h *DepartmentHandler) GetDepartment(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)

	depth := 1
	if d := r.URL.Query().Get("depth"); d != "" {
		if parsed, err := strconv.Atoi(d); err == nil && parsed > 0 && parsed <= 5 {
			depth = parsed
		}
	}

	includeEmp := true
	if q := r.URL.Query().Get("include_employees"); q == "false" {
		includeEmp = false
	}

	dep, err := h.service.GetDepartmentTree(uint(id), depth, includeEmp)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	json.NewEncoder(w).Encode(dep)
}

func (h *DepartmentHandler) UpdateDepartment(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)

	var req struct {
		Name     *string `json:"name"`
		ParentID *uint   `json:"parent_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", 400)
		return
	}

	dep, err := h.service.UpdateDepartment(uint(id), req.Name, req.ParentID)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(dep)
}

func (h *DepartmentHandler) DeleteDepartment(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)

	mode := r.URL.Query().Get("mode")
	reassignStr := r.URL.Query().Get("reassign_to_department_id")
	var reassignID *uint
	if reassignStr != "" {
		if parsed, err := strconv.Atoi(reassignStr); err == nil {
			tmp := uint(parsed)
			reassignID = &tmp
		}
	}

	if err := h.service.DeleteDepartment(uint(id), mode, reassignID); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
