package routes

import (
	"time"

	"github.com/JamesCCoder/resume_backend_go/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()

	router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

    api := router.Group("")
    {
        api.GET("/students", controllers.GetStudents)
        api.POST("/students", controllers.CreateStudent)
        api.GET("/students/:id", controllers.GetStudent)
        api.PUT("/students/:id", controllers.UpdateStudent)
        api.DELETE("/students/:id", controllers.DeleteStudent)

        api.GET("/professors", controllers.GetProfessors)
        api.POST("/professors", controllers.CreateProfessor)
        api.GET("/professors/:id", controllers.GetProfessor)
        api.PUT("/professors/:id", controllers.UpdateProfessor)
        api.DELETE("/professors/:id", controllers.DeleteProfessor)

        api.POST("api/login", controllers.AdminLogin)
    }


    return router
}
