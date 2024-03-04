package http

import (
	"crud/internal/core/interface/service"
	"crud/internal/transport/handler"
	"crud/internal/transport/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoutes(service service.AuthService, postService service.PostService, commentService service.CommentService) *gin.Engine {
	router := gin.New()

	router.POST("/register", handler.RegisterUser(service))

	api := router.Group("/api", middleware.AuthMiddleware)
	{
		api.POST("/post", handler.CreatePost(postService))
		api.GET("/post/:id", handler.GetPost(postService))
		api.DELETE("/delete/:id", handler.DeletePost(postService))
		api.PUT("/update/:id", handler.UpdatePost(postService))
	}

	com := router.Group("/comment", middleware.AuthMiddleware)
	{
		com.POST("/comment", handler.CreateComment(commentService))
		com.GET("/comment/:id", handler.GetComment(commentService))
		com.DELETE("/delete/:id", handler.DeleteComment(commentService))
		com.PUT("/update/:id", handler.UpdateComment(commentService))
	}

	return router
}
