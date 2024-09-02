package handler

import (
	"context"
	"log"

	author_proto "author-service/authorpb"
	"author-service/internal/interfaces"
)

type AuthorHandler struct {
	author_proto.UnimplementedAuthorServiceServer
	repoAuthor interfaces.AuthorInterfaceHandler
}

func NewHandler(repoAuthor interfaces.AuthorInterfaceHandler) *AuthorHandler {
	return &AuthorHandler{repoAuthor: repoAuthor}
}

func (uc *AuthorHandler) GetAllAuthors(ctx context.Context, payload *author_proto.GetAllAuthorsRequest) (*author_proto.GetAllAuthorsResponse, error) {
	result, err := uc.repoAuthor.GetAllAuthors(ctx, payload)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (uc *AuthorHandler) CreateAuthor(ctx context.Context, payload *author_proto.CreateAuthorRequest) (*author_proto.Author, error) {

	result, err := uc.repoAuthor.CreateAuthor(ctx, payload)

	if err != nil {
		return nil, err
	}

	return result, nil

}

func (uc *AuthorHandler) GetAuthorById(ctx context.Context, payload *author_proto.GetDetailAuthorRequest) (*author_proto.GetAuthorResponse, error) {

	result, err := uc.repoAuthor.GetAuthorById(ctx, payload)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil

}

func (uc *AuthorHandler) UpdateAuthor(ctx context.Context, payload *author_proto.UpdateAuthorRequest) (*author_proto.Author, error) {

	result, err := uc.repoAuthor.UpdateAuthor(ctx, payload)

	if err != nil {
		return nil, err
	}

	return result, nil

}

func (uc *AuthorHandler) DeleteAuthor(ctx context.Context, id *author_proto.DeleteAuthorRequest) (*author_proto.DeleteAuthorResponse, error) {

	result, err := uc.repoAuthor.DeleteAuthor(ctx, id)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil

}
