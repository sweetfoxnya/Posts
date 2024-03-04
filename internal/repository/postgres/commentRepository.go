package postgres

import (
	"context"
	"crud/internal/core/interface/repository"
	"crud/internal/core/model"
	"crud/internal/lib/db"
	"crud/internal/repository/dbModel"
	"fmt"
)

type _commentRepository struct {
	db *db.Db
}

func NewCommentRepo(db *db.Db) repository.CommentRepository {
	return _commentRepository{db}
}

func (commentRepository _commentRepository) CreateComment(ctx context.Context, comment model.Comment) (int, error) {
	commentDb := dbModel.Comment(comment)
	var id int

	err := commentRepository.db.PgConn.QueryRow(ctx,
		`INSERT INTO public.comment(text, userID, postID) values ($1,$2,$3) RETURNING id`,
		commentDb.Text,
		commentDb.UserID,
		commentDb.PostID).Scan(&id)

	return id, err
}

func (commentRepository _commentRepository) GetComment(ctx context.Context, commentId int) (model.Comment, error) {
	var comment dbModel.Comment

	err := commentRepository.db.PgConn.QueryRow(ctx,
		`SELECT c.text, c.userID, c.postID FROM public.comment c WHERE c.id=$1`,
		commentId).Scan(&comment.Text, &comment.UserID, &comment.PostID)

	if err != nil {
		return model.Comment{}, fmt.Errorf("ошибка получения комментария: %s", err.Error())
	}

	return model.Comment(comment), nil

}

func (commentRepository _commentRepository) DeleteComment(ctx context.Context, commentId int) error {
	_, err := commentRepository.db.PgConn.Exec(ctx, `DELETE FROM public.comment WHERE id = $1`, commentId)
	if err != nil {
		return fmt.Errorf("ошибка удаления комментария: %s", err.Error())
	}

	return nil
}

func (commentRepository _commentRepository) UpdateComment(ctx context.Context, comment model.Comment, commentId int) error {
	commentDb := dbModel.Comment(comment)

	_, err := commentRepository.db.PgConn.Exec(ctx,
		`UPDATE public.comment SET text=$1, userID=$2, postID=$3 WHERE id=$4`,
		commentDb.Text,
		commentDb.UserID,
		commentDb.PostID,
		commentId)

	if err != nil {
		return fmt.Errorf("ошибка изменения комментария: %s", err.Error())
	}

	return nil
}
