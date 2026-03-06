package service

import (
	"errors"
	"strings"
	"time"

	"org-structure-api/internal/models"
	"org-structure-api/internal/repository"

	"gorm.io/gorm"
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

	dep := &models.Department{Name: name, ParentID: parentID, CreatedAt: time.Now()}
	if err := s.repo.CreateDepartment(dep); err != nil {
		return nil, err
	}
	return dep, nil
}

func (s *DepartmentService) CreateEmployee(depID uint, fullName, position string, hiredAt *string) (*models.Employee, error) {
	fullName = strings.TrimSpace(fullName)
	position = strings.TrimSpace(position)
	if fullName == "" || position == "" {
		return nil, errors.New("invalid data")
	}

	_, err := s.repo.GetDepartment(depID)
	if err != nil {
		return nil, errors.New("department not found")
	}

	var hired *time.Time
	if hiredAt != nil {
		t, err := time.Parse("2006-01-02", *hiredAt)
		if err != nil {
			return nil, errors.New("invalid hired_at date format, expected YYYY-MM-DD")
		}
		hired = &t
	}

	emp := &models.Employee{
		DepartmentID: depID,
		FullName:     fullName,
		Position:     position,
		HiredAt:      hired,
		CreatedAt:    time.Now(),
	}

	if err := s.repo.CreateEmployee(emp); err != nil {
		return nil, err
	}

	return emp, nil
}

func (s *DepartmentService) GetDepartmentTree(id uint, depth int, includeEmployees bool) (*models.Department, error) {
	dep, err := s.repo.GetDepartment(id)
	if err != nil {
		return nil, err
	}

	if includeEmployees {
		emps, _ := s.repo.GetEmployees(id)
		dep.Employees = emps
	}

	if depth > 0 {
		children, _ := s.repo.GetChildren(id)
		for i := range children {
			child, _ := s.GetDepartmentTree(children[i].ID, depth-1, includeEmployees)
			children[i] = *child
		}
		dep.Children = children
	}

	return dep, nil
}

func (s *DepartmentService) UpdateDepartment(id uint, name *string, parentID *uint) (*models.Department, error) {
	dep, err := s.repo.GetDepartment(id)
	if err != nil {
		return nil, errors.New("department not found")
	}

	if name != nil {
		n := strings.TrimSpace(*name)
		if n == "" {
			return nil, errors.New("name cannot be empty")
		}
		dep.Name = n
	}

	if parentID != nil {
		if *parentID == dep.ID {
			return nil, errors.New("cannot set department as its own parent")
		}
		// TODO: проверить на циклы (перемещение в поддерево)
		dep.ParentID = parentID
	}

	if err := s.repo.UpdateDepartment(dep); err != nil {
		return nil, err
	}

	return dep, nil
}

func (s *DepartmentService) DeleteDepartment(id uint, mode string, reassignID *uint) error {
	_, err := s.repo.GetDepartment(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("department not found")
		}
		return err
	}

	switch mode {
	case "cascade":
		children, _ := s.repo.GetChildren(id)
		for _, child := range children {
			s.DeleteDepartment(child.ID, "cascade", nil)
		}
		emps, _ := s.repo.GetEmployees(id)
		for _, emp := range emps {
			s.repo.DeleteEmployee(emp.ID)
		}
		return s.repo.DeleteDepartment(id)
	case "reassign":
		if reassignID == nil {
			return errors.New("reassign_to_department_id required")
		}
		emps, _ := s.repo.GetEmployees(id)
		for _, emp := range emps {
			emp.DepartmentID = *reassignID
			s.repo.UpdateEmployee(&emp)
		}
		children, _ := s.repo.GetChildren(id)
		for _, child := range children {
			s.DeleteDepartment(child.ID, "reassign", reassignID)
		}
		return s.repo.DeleteDepartment(id)
	default:
		return errors.New("invalid mode, must be 'cascade' or 'reassign'")
	}
}
