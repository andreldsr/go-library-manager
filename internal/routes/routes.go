package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-library-manager/internal/controllers"
	"go-library-manager/internal/middlewares"
	"os"
	"strings"
	"time"
)

func HandleRequests() {
	r := gin.Default()
	err := r.SetTrustedProxies(nil)
	if err != nil {
		return
	}

	r.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Split(os.Getenv("CORS_ORIGINS"), ","),
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/api/book", controllers.FindAllBooks)
	r.GET("/api/book/:id", controllers.FindBookById)
	r.GET("/api/book/stats", controllers.GetBookStats)
	r.POST("/api/auth", controllers.Login)
	r.Use(middlewares.JWTAuthMiddleware())

	r.POST("/api/auth/register", controllers.Register)

	adminRoutes := r.Group("/api/admin")
	adminRoutes.Use(middlewares.AdminMiddleware())

	adminRoutes.POST("/book", controllers.CreateBook)

	adminRoutes.GET("/lending/open", controllers.FindAllLendingsActive)
	adminRoutes.GET("/lending/today", controllers.FindAllLendingsDueToday)
	adminRoutes.GET("/lending/overdue", controllers.FindAllLendingsOverdue)
	adminRoutes.GET("/lending/:id", controllers.FindLendingById)
	adminRoutes.POST("/lending", controllers.CreateLending)
	adminRoutes.PATCH("/lending/:id/return", controllers.ReturnLending)

	adminRoutes.GET("/user", controllers.FindUserList)
	adminRoutes.GET("/user/:id", controllers.FindUserById)
	adminRoutes.PUT("/user/:id", controllers.UpdateUser)
	_ = r.Run()
}
