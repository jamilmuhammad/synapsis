package interfaces

import (
	"context"
	"sharing_vasion_indonesia/api_gateway/internal/models"
	article_proto "sharing_vasion_indonesia/pkg/proto"
)

type ArticleUseCase interface {
	GetArticles(ctx context.Context, limit string, offset string) (*article_proto.GetArticlesResponse, error)
	CreateArticle(ctx context.Context, payload models.CreateArticleRequest) (*article_proto.Post, error)
	UpdateArticle(ctx context.Context, payload models.UpdateArticleRequest) (*article_proto.Post, error)
	GetArticle(ctx context.Context, id string) (*article_proto.GetArticleResponse, error)
	DeleteArticle(ctx context.Context, id string) error
}
