package handler

import (
	"crud/internal/core/interface/service"
	"crud/internal/core/model"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

type handlerPost struct {
	Title    string `json:"title"`
	Body     string `json:"body"`
	ImageURL string `json:"image"`
	Author   string `json:"author"`
}

func CreatePost(service service.PostService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var post handlerPost

		login := c.GetString("user")

		if err := c.BindJSON(&post); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "неверное тело запроса"})

			return
		}

		post.Author = login

		id, err := service.CreatePost(c.Request.Context(), model.Post(post))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"post": id})
	}
}

func GetPost(service service.PostService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		numberId, err := strconv.Atoi(id)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "неверно передан id поста"})

			return
		}

		post, err := service.GetPost(c.Request.Context(), numberId)

		if err != nil {
			slog.Error(err.Error())

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "ошибка получения поста"})

			return

		}

		c.JSON(http.StatusOK, handlerPost(post))

	}
}

func DeletePost(service service.PostService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		numberId, err := strconv.Atoi(id)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "неверно передан id поста"})

			return
		}

		err = service.DeletePost(c.Request.Context(), numberId)

		if err != nil {
			slog.Error(err.Error())

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "ошибка удаления поста"})

			return

		}

		c.JSON(http.StatusOK, gin.H{"message": "пост успешно удален"})

	}
}

func UpdatePost(service service.PostService) gin.HandlerFunc {
	return func(c *gin.Context) {

		var updatedPost model.Post

		id := c.Param("id")

		numberId, err := strconv.Atoi(id)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "неверно передан id поста"})

			return
		}

		if err := c.BindJSON(&updatedPost); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "неверное тело запроса"})

			return
		}

		err = service.UpdatePost(c.Request.Context(), updatedPost, numberId)

		if err != nil {
			slog.Error(err.Error())

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "ошибка изменения поста"})

			return

		}

		c.JSON(http.StatusOK, gin.H{"message": "пост успешно изменен"})

	}
}
