package router

import (
	"gin/basic/controller"

	"github.com/gin-gonic/gin"
)

func UserRouter(engine *gin.Engine) {
	engine.POST("/signup", controller.SignUp) //getPost this will come from controller
	engine.POST("/login", controller.LogIn)
	engine.POST("/logout", controller.LogOut)
	engine.PUT("/forgetpassword", controller.ForgetPassword)
	//engine.PUT("/post", controller.GetPost)
	//engine.DELETE("/user", controller.DeletePost)
	//engine.PATCH("/post", controller.GetPost)
}
