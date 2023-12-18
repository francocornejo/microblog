package api

import (
	"microblog/handler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(r *gin.Engine, handler handler.MicroBlogHandler) {
	v1 := r.Group("/microblog")
	{
		v1.POST("/send", handler.SendMessageHandler)
		v1.POST("/follow", handler.FollowHandler)
		v1.GET("/messages", handler.TimelineHandler)
	}
	r.GET("/api/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // ../api/doc/index.html
}
