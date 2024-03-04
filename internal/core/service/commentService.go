package service

import (
	"context"
	"crud/internal/core/interface/repository"
	"crud/internal/core/interface/service"
	"crud/internal/core/model"
	"errors"
	"log/slog"
)

type _commentService struct {
	repo repository.CommentRepository
}

func NewCommentService(repo repository.CommentRepository) service.CommentService {
	return _commentService{repo: repo}
}

func (commentService _commentService) CreateComment(ctx context.Context, comment model.Comment) (int, error) {
	id, err := commentService.repo.CreateComment(ctx, comment)

	if err != nil {
		slog.Error(err.Error())
		return 0, errors.New("ошибка создания комментария")
	}

	return id, nil
}

func (commentService _commentService) GetComment(ctx context.Context, commentId int) (model.Comment, error) {
	return commentService.repo.GetComment(ctx, commentId)
}

func (commentService _commentService) DeleteComment(ctx context.Context, commentId int) error {
	err := commentService.repo.DeleteComment(ctx, commentId)

	if err != nil {
		slog.Error(err.Error())
		return errors.New("ошибка удаления комметария")
	}

	return nil
}

func (commentService _commentService) UpdateComment(ctx context.Context, comment model.Comment, commentId int) error {
	err := commentService.repo.UpdateComment(ctx, comment, commentId)

	if err != nil {
		slog.Error(err.Error())
		return errors.New("ошибка изменения комментария")
	}

	return nil
}
