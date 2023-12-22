package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-library-manager/internal/controllers"
	"go-library-manager/internal/middlewares"
	"os"
	"strings"
)

func HandleRequests() {
	r := gin.Default()
	err := r.SetTrustedProxies(nil)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Split(os.Getenv("CORS_ORIGINS"), ","),
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
	}))
	if err != nil {
		return
	}
	r.POST("/api/auth", controllers.Login)
	r.Use(middlewares.JWTAuthMiddleware())

	r.POST("/api/auth/register", controllers.Register)
	r.GET("/api/book", controllers.FindAllBooks)
	r.GET("/api/book/:id", controllers.FindBookById)
	r.GET("/api/book/stats", controllers.GetBookStats)

	adminRoutes := r.Group("/api/admin")
	adminRoutes.Use(middlewares.AdminMiddleware())

	adminRoutes.GET("/lending/open", controllers.FindAllLendingsActive)
	adminRoutes.GET("/lending/today", controllers.FindAllLendingsDueToday)
	adminRoutes.GET("/lending/overdue", controllers.FindAllLendingsOverdue)
	adminRoutes.GET("/user/name/:name", controllers.FindUserList)
	adminRoutes.GET("/lending/:id", controllers.FindLendingById)
	adminRoutes.POST("/lending", controllers.CreateLending)
	adminRoutes.PATCH("/lending/:id/return", controllers.ReturnLending)
	adminRoutes.POST("/book", controllers.CreateBook)
	_ = r.Run()
}
