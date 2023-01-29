package router

import (
	"gin/basic/controller"

	"github.com/gin-gonic/gin"
)

func UserRouter(engine *gin.Engine) {
	engine.GET("/user", controller.GetPost) //getPost this will come from controller
	engine.POST("/user", controller.CreatePost)

	//engine.PUT("/post", controller.GetPost)
	//engine.DELETE("/user", controller.DeletePost)
	//engine.PATCH("/post", controller.GetPost)
}
