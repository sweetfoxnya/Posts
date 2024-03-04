package service

import (
	"context"
	"crud/internal/core/model"
)

type AuthService interface {
	Register(ctx context.Context, login, password string) (string, error)
	GenerateToken(ctx context.Context, login, password string) (string, error)
}

type PostService interface {
	CreatePost(ctx context.Context, post model.Post) (int, error)
	GetPost(ctx context.Context, postId int) (model.Post, error)
	DeletePost(ctx context.Context, postId int) error
	UpdatePost(ctx context.Context, post model.Post, postId int) error
}

type CommentService interface {
	CreateComment(ctx context.Context, comment model.Comment) (int, error)
	GetComment(ctx context.Context, commentId int) (model.Comment, error)
	DeleteComment(ctx context.Context, commentId int) error
	UpdateComment(ctx context.Context, comment model.Comment, commentId int) error
}
