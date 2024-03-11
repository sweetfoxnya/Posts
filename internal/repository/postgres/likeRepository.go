package postgres

import (
	"context"
	"crud/internal/core/interface/repository"
	"crud/internal/core/model"
	"crud/internal/lib/db"
	"crud/internal/repository/dbModel"
	"fmt"
)

type _likeRepository struct {
	db *db.Db
}

func NewLikeRepo(db *db.Db) repository.LikeRepository {
	return _likeRepository{db}
}

func (likeRepository _likeRepository) PutLike(ctx context.Context, like model.Like) (int, error) {
	likeDb := dbModel.Like(like)
	var id int

	err := likeRepository.db.PgConn.QueryRow(ctx,
		`INSERT INTO public.like(postID) values ($1) RETURNING id`,
		likeDb.PostID).Scan(&id)

	return id, err
}

func (likeRepository _likeRepository) GetLike(ctx context.Context, likeId int) (model.Like, error) {
	var like dbModel.Like

	err := likeRepository.db.PgConn.QueryRow(ctx,
		`SELECT l.postID FROM public.like l WHERE l.id=$1`,
		likeId).Scan(&like.PostID)

	if err != nil {
		return model.Like{}, fmt.Errorf("ошибка получения лайков: %s", err.Error())
	}

	return model.Like(like), err

}

func (likeRepository _likeRepository) DeleteLike(ctx context.Context, likeId int) error {
	_, err := likeRepository.db.PgConn.Exec(ctx, `DELETE FROM public.like WHERE id = $1`, likeId)
	if err != nil {
		return fmt.Errorf("ошибка удаления лайка: %s", err.Error())
	}

	return err
}
