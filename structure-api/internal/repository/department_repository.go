package repository

import (
	"org-structure-api/internal/models"

	"gorm.io/gorm"
)

type DepartmentRepository struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) *DepartmentRepository {
	return &DepartmentRepository{db: db}
}

// Department methods
func (r *DepartmentRepository) CreateDepartment(dep *models.Department) error {
	return r.db.Create(dep).Error
}

func (r *DepartmentRepository) GetDepartment(id uint) (*models.Department, error) {
	var dep models.Department
	err := r.db.First(&dep, id).Error
	return &dep, err
}

func (r *DepartmentRepository) GetChildren(depID uint) ([]models.Department, error) {
	var deps []models.Department
	err := r.db.Where("parent_id = ?", depID).Find(&deps).Error
	return deps, err
}

func (r *DepartmentRepository) UpdateDepartment(dep *models.Department) error {
	return r.db.Save(dep).Error
}

func (r *DepartmentRepository) DeleteDepartment(id uint) error {
	return r.db.Delete(&models.Department{}, id).Error
}

// Employee methods
func (r *DepartmentRepository) CreateEmployee(emp *models.Employee) error {
	return r.db.Create(emp).Error
}

func (r *DepartmentRepository) GetEmployees(depID uint) ([]models.Employee, error) {
	var employees []models.Employee
	err := r.db.Where("department_id = ?", depID).
		Order("full_name").
		Find(&employees).Error
	return employees, err
}

func (r *DepartmentRepository) UpdateEmployee(emp *models.Employee) error {
	return r.db.Save(emp).Error
}

func (r *DepartmentRepository) DeleteEmployee(id uint) error {
	return r.db.Delete(&models.Employee{}, id).Error
}
