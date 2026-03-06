package models

import "time"

type Department struct {
	ID        uint         `gorm:"primaryKey" json:"id"`
	Name      string       `json:"name"`
	ParentID  *uint        `json:"parent_id"`
	CreatedAt time.Time    `json:"created_at"`
	Employees []Employee   `json:"employees,omitempty"`
	Children  []Department `gorm:"foreignKey:ParentID" json:"children,omitempty"`
}

type Employee struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	DepartmentID uint       `json:"department_id"`
	FullName     string     `json:"full_name"`
	Position     string     `json:"position"`
	HiredAt      *time.Time `json:"hired_at"`
	CreatedAt    time.Time  `json:"created_at"`
}
