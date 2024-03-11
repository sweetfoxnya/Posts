package repository

import (
	"context"
	"crud/internal/core/model"
)

type AuthRepository interface {
	GetUser(ctx context.Context, login, hashPassword string) (string, error)
	Register(ctx context.Context, login, hashPassword string) (string, error)
}

type PostRepository interface {
	CreatePost(ctx context.Context, post model.Post) (int, error)
	GetPost(ctx context.Context, postId int) (model.Post, error)
	DeletePost(ctx context.Context, postId int) error
	UpdatePost(ctx context.Context, post model.Post, postId int) error
}

type CommentRepository interface {
	CreateComment(ctx context.Context, comment model.Comment) (int, error)
	GetComment(ctx context.Context, commentId int) (model.Comment, error)
	DeleteComment(ctx context.Context, commentId int) error
	UpdateComment(ctx context.Context, comment model.Comment, commentId int) error
}

type LikeRepository interface {
	PutLike(ctx context.Context, like model.Like) (int, error)
	GetLike(ctx context.Context, likeId int) (model.Like, error)
	DeleteLike(ctx context.Context, likeId int) error
}
