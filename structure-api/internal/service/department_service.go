package service

import (
	"errors"
	"strings"

	"org-structure-api/internal/models"
	"org-structure-api/internal/repository"
)

type DepartmentService struct {
	repo *repository.DepartmentRepository
}

func NewDepartmentService(r *repository.DepartmentRepository) *DepartmentService {
	return &DepartmentService{repo: r}
}

func (s *DepartmentService) CreateDepartment(name string, parentID *uint) (*models.Department, error) {

	name = strings.TrimSpace(name)
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	dep := &models.Department{
		Name:     name,
		ParentID: parentID,
	}

	err := s.repo.CreateDepartment(dep)
	return dep, err
}

func (s *DepartmentService) CreateEmployee(depID uint, fullName, position string) (*models.Employee, error) {

	fullName = strings.TrimSpace(fullName)
	position = strings.TrimSpace(position)

	if fullName == "" || position == "" {
		return nil, errors.New("invalid data")
	}

	_, err := s.repo.GetDepartment(depID)
	if err != nil {
		return nil, err
	}

	emp := &models.Employee{
		DepartmentID: depID,
		FullName:     fullName,
		Position:     position,
	}

	err = s.repo.CreateEmployee(emp)

	return emp, err
}

func (s *DepartmentService) GetDepartmentTree(id uint, depth int) (*models.Department, error) {

	dep, err := s.repo.GetDepartment(id)
	if err != nil {
		return nil, err
	}

	if depth > 0 {
		children, _ := s.repo.GetChildren(id)

		for i := range children {
			child, _ := s.GetDepartmentTree(children[i].ID, depth-1)
			children[i] = *child
		}

		dep.Children = children
	}

	emps, _ := s.repo.GetEmployees(id)
	dep.Employees = emps

	return dep, nil
}
