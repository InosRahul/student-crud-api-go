package main

import (
	"github.com/InosRahul/student-crud-api/config"
	controller "github.com/InosRahul/student-crud-api/controllers"
	"github.com/InosRahul/student-crud-api/models"
	"github.com/InosRahul/student-crud-api/repository"
	service "github.com/InosRahul/student-crud-api/services"
	"github.com/InosRahul/student-crud-api/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
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

	repo := repository.NewStudentRepository(db)
	service := service.NewStudentService(repo)
	controller := controller.NewStudentController(service)

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.POST("/students", controller.CreateStudent)
		v1.GET("/students", controller.GetAllStudents)
		v1.GET("/students/:id", controller.GetStudentByID)
		v1.PUT("/students/:id", controller.UpdateStudent)
		v1.DELETE("/students/:id", controller.DeleteStudent)
	}

	router.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "OK"})
	})

	utils.Logger.Info("Starting server on port ", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		utils.Logger.Fatal("Failed to start server: ", err)
	}
}
