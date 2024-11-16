package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/InosRahul/student-crud-api/config"
	"github.com/InosRahul/student-crud-api/models"
	"github.com/InosRahul/student-crud-api/repository"
	service "github.com/InosRahul/student-crud-api/services"
	"github.com/InosRahul/student-crud-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	cfg, err := config.LoadConfig()
	if err != nil {
		utils.Logger.Fatal("Failed to load config: ", err)
	}

	dsn := cfg.GetDBConnectionString()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		utils.Logger.Fatal("Failed to connect to database: ", err)
	}

	// Auto migrate the schema
	db.AutoMigrate(&models.Student{})
	return db
}

func TestStudentController(t *testing.T) {
	db := setupTestDB()
	repo := repository.NewStudentRepository(db)
	service := service.NewStudentService(repo)
	controller := NewStudentController(service)

	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.POST("/students", controller.CreateStudent)
		v1.GET("/students", controller.GetAllStudents)
		v1.GET("/students/:id", controller.GetStudentByID)
		v1.PUT("/students/:id", controller.UpdateStudent)
		v1.DELETE("/students/:id", controller.DeleteStudent)
	}

	t.Run("CreateStudent", func(t *testing.T) {
		student := models.Student{Name: "John Doe", Email: "john@example.com", Age: 20, Course: "C1"}
		body, _ := json.Marshal(student)
		req, _ := http.NewRequest("POST", "/api/v1/students", bytes.NewBuffer(body))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		var responseStudent models.Student
		json.Unmarshal(w.Body.Bytes(), &responseStudent)
		assert.Equal(t, student.Name, responseStudent.Name)
		assert.Equal(t, student.Email, responseStudent.Email)
		assert.Equal(t, student.Age, responseStudent.Age)
		assert.Equal(t, student.Course, responseStudent.Course)
	})

	t.Run("GetAllStudents", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/students", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		var students []models.Student
		json.Unmarshal(w.Body.Bytes(), &students)
		assert.NotEmpty(t, students)
	})

	t.Run("GetStudentByID", func(t *testing.T) {
		student := models.Student{Name: "Jane Doe", Email: "jane@example.com"}
		db.Create(&student)

		req, _ := http.NewRequest("GET", "/api/v1/students/"+strconv.Itoa(student.ID), nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		var responseStudent models.Student
		json.Unmarshal(w.Body.Bytes(), &responseStudent)
		assert.Equal(t, student.Name, responseStudent.Name)
		assert.Equal(t, student.Email, responseStudent.Email)
	})

	t.Run("UpdateStudent", func(t *testing.T) {
		student := models.Student{Name: "John Doe", Email: "john@example.com", Age: 24, Course: "C1"}
		db.Create(&student)

		updatedStudent := models.Student{Name: "John Updated", Email: "john.updated@example.com", Age: 24, Course: "C1"}
		body, _ := json.Marshal(updatedStudent)
		req, _ := http.NewRequest("PUT", "/api/v1/students/"+strconv.Itoa(student.ID), bytes.NewBuffer(body))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		var responseStudent models.Student
		json.Unmarshal(w.Body.Bytes(), &responseStudent)
		assert.Equal(t, updatedStudent.Name, responseStudent.Name)
		assert.Equal(t, updatedStudent.Email, responseStudent.Email)
	})

	t.Run("DeleteStudent", func(t *testing.T) {
		student := models.Student{Name: "John Doe", Email: "john@example.com"}
		db.Create(&student)

		req, _ := http.NewRequest("DELETE", "/api/v1/students/"+strconv.Itoa(student.ID), nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		var response map[string]string
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "Student deleted", response["message"])
	})
}
