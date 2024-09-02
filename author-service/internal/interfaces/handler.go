package interfaces

import (
	author_proto "author-service/authorpb"
	"context"
)

type AuthorInterfaceHandler interface {
	GetAllAuthors(ctx context.Context, payload *author_proto.GetAllAuthorsRequest) (*author_proto.GetAllAuthorsResponse, error)
	CreateAuthor(ctx context.Context, post *author_proto.CreateAuthorRequest) (*author_proto.Author, error)
	UpdateAuthor(ctx context.Context, post *author_proto.UpdateAuthorRequest) (*author_proto.Author, error)
	GetAuthorById(ctx context.Context, payload *author_proto.GetDetailAuthorRequest) (*author_proto.GetAuthorResponse, error)
	DeleteAuthor(ctx context.Context, id *author_proto.DeleteAuthorRequest) (*author_proto.DeleteAuthorResponse, error)
}
