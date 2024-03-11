package handler

import (
	"crud/internal/core/interface/service"
	"crud/internal/core/model"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

type handlerLike struct {
	PostID int `json:"postID"`
}

func PutLike(service service.LikeService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var like handlerLike

		if err := c.BindJSON(&like); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "неверное тело запроса"})

			return
		}

		id, err := service.PutLike(c.Request.Context(), model.Like(like))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"like": id})
	}
}

func GetLikes(service service.LikeService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		numberId, err := strconv.Atoi(id)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "неверно передан id поста"})

			return
		}

		like, err := service.GetLike(c.Request.Context(), numberId)

		if err != nil {
			slog.Error(err.Error())

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "ошибка получения лайков"})

			return

		}

		c.JSON(http.StatusOK, handlerLike(like))

	}
}

func DeleteLike(service service.LikeService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		numberId, err := strconv.Atoi(id)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "неверно передан id лайка"})

			return
		}

		err = service.DeleteLike(c.Request.Context(), numberId)

		if err != nil {
			slog.Error(err.Error())

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "ошибка удаления лайка"})

			return

		}

		c.JSON(http.StatusOK, gin.H{"message": "лайк успешно удален"})

	}
}
