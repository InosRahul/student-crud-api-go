package controller

import (
	"net/http"
	"strconv"

	model "github.com/InosRahul/student-crud-api/models"
	service "github.com/InosRahul/student-crud-api/services"
	"github.com/gin-gonic/gin"
)

type StudentController struct {
	service *service.StudentService
}

func NewStudentController(service *service.StudentService) *StudentController {
	return &StudentController{service: service}
}

func (c *StudentController) CreateStudent(ctx *gin.Context) {
	var student model.Student
	if err := ctx.ShouldBindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.service.CreateStudent(&student); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, student)
}

func (c *StudentController) GetAllStudents(ctx *gin.Context) {
	students, err := c.service.GetAllStudents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, students)
}

func (c *StudentController) GetStudentByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	student, err := c.service.GetStudentByID(int(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	ctx.JSON(http.StatusOK, student)
}

func (c *StudentController) UpdateStudent(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var student model.Student
	if err := ctx.ShouldBindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	student.ID = int(id)
	if err := c.service.UpdateStudent(&student); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, student)
}

func (c *StudentController) DeleteStudent(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := c.service.DeleteStudent(int(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Student deleted"})
}
