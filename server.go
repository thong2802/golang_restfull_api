package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thong2802/golang_api/config"
	"github.com/thong2802/golang_api/controller"
	"gorm.io/gorm"
)

var(
	db				   			*gorm.DB 	 = config.SetUpDatabaseConnection()
	authController controller.AuthController = controller.NewAuthController()
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	authRouters := r.Group("api/auth") 
	{
		authRouters.POST("/login", controller.NewAuthController().Login)
		authRouters.POST("/register", controller.NewAuthController().Register)
	}
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
