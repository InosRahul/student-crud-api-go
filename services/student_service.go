package service

import (
	"github.com/InosRahul/student-crud-api/models"
	"github.com/InosRahul/student-crud-api/repository"
)

type StudentService struct {
	repo *repository.StudentRepository
}

func NewStudentService(repo *repository.StudentRepository) *StudentService {
	return &StudentService{repo: repo}
}

func (s *StudentService) CreateStudent(student *models.Student) error {
	return s.repo.CreateStudent(student)
}

func (s *StudentService) GetAllStudents() ([]models.Student, error) {
	return s.repo.GetAllStudents()
}

func (s *StudentService) GetStudentByID(id int) (*models.Student, error) {
	return s.repo.GetStudentByID(uint(id))
}

func (s *StudentService) UpdateStudent(student *models.Student) error {
	return s.repo.UpdateStudent(student)
}

func (s *StudentService) DeleteStudent(id int) error {
	return s.repo.DeleteStudent(uint(id))
}
