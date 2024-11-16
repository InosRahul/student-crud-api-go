package repository

import (
	"github.com/InosRahul/student-crud-api/models"
	"gorm.io/gorm"
)

type StudentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) *StudentRepository {
	return &StudentRepository{db: db}
}

func (r *StudentRepository) CreateStudent(student *models.Student) error {
	return r.db.Create(student).Error
}

func (r *StudentRepository) GetAllStudents() ([]models.Student, error) {
	var students []models.Student
	err := r.db.Find(&students).Error
	return students, err
}

func (r *StudentRepository) GetStudentByID(id uint) (*models.Student, error) {
	var student models.Student
	err := r.db.First(&student, id).Error
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *StudentRepository) UpdateStudent(student *models.Student) error {
	return r.db.Save(student).Error
}

func (r *StudentRepository) DeleteStudent(id uint) error {
	return r.db.Delete(&models.Student{}, id).Error
}
