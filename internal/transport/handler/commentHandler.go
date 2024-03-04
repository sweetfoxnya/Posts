package handler

import (
	"crud/internal/core/interface/service"
	"crud/internal/core/model"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

type handlerComment struct {
	Text   string `json:"text"`
	User   string `json:"user"`
	PostID int    `json:"postID"`
}

func CreateComment(service service.CommentService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var comment handlerComment

		login := c.GetString("user")

		if err := c.BindJSON(&comment); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "неверное тело запроса"})

			return
		}

		comment.User = login

		id, err := service.CreateComment(c.Request.Context(), model.Comment(comment))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"comment": id})
	}
}

func GetComment(service service.CommentService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		numberId, err := strconv.Atoi(id)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "неверно передан id комментария"})

			return
		}

		comment, err := service.GetComment(c.Request.Context(), numberId)

		if err != nil {
			slog.Error(err.Error())

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "ошибка получения комментария"})

			return

		}

		c.JSON(http.StatusOK, handlerComment(comment))

	}
}

func DeleteComment(service service.CommentService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		numberId, err := strconv.Atoi(id)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "неверно передан id комментария"})

			return
		}

		err = service.DeleteComment(c.Request.Context(), numberId)

		if err != nil {
			slog.Error(err.Error())

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "ошибка удаления комментария"})

			return

		}

		c.JSON(http.StatusOK, gin.H{"message": "комментарий успешно удален"})

	}
}

func UpdateComment(service service.CommentService) gin.HandlerFunc {
	return func(c *gin.Context) {

		var updatedComment model.Comment

		id := c.Param("id")

		numberId, err := strconv.Atoi(id)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "неверно передан id комменатрия"})

			return
		}

		err = service.UpdateComment(c.Request.Context(), updatedComment, numberId)

		if err != nil {
			slog.Error(err.Error())

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "ошибка изменения комментария"})

			return

		}

		c.JSON(http.StatusOK, gin.H{"message": "комментарий успешно изменен"})

	}
}
