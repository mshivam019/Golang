package router

import (
    "gin/basic/controller"
    "github.com/gin-gonic/gin"
    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/files"
 _	"gin/basic/docs"
)

func UserRouter(engine *gin.Engine) {
    engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    engine.POST("/signup", controller.SignUp)
    engine.POST("/login", controller.LogIn)
    engine.POST("/logout", controller.LogOut)
}
