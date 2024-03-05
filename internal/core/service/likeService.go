package service

import (
	"context"
	"crud/internal/core/interface/repository"
	"crud/internal/core/interface/service"
	"crud/internal/core/model"
	"errors"
	"log/slog"
)

type _likeService struct {
	repo repository.LikeRepository
}

func NewLikeService(repo repository.LikeRepository) service.LikeService {
	return _likeService{repo: repo}
}

func (likeService _likeService) PutLike(ctx context.Context, like model.Like) (int, error) {
	id, err := likeService.repo.PutLike(ctx, like)

	if err != nil {
		slog.Error(err.Error())
		return 0, errors.New("ошибка постановки лайка")
	}

	return id, nil
}

func (likeService _likeService) GetLikes(ctx context.Context, postId int) (model.Like, error) {
	return likeService.repo.GetLikes(ctx, postId)
}

func (likeService _likeService) DeleteLike(ctx context.Context, likeId int) error {
	err := likeService.repo.DeleteLike(ctx, likeId)

	if err != nil {
		slog.Error(err.Error())
		return errors.New("ошибка удаления лайка")
	}

	return nil
}
