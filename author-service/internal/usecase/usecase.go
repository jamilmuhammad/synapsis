package usecase

import (
	"context"
	"lib"
	"log"

	author_proto "author-service/authorpb"
	"author-service/internal/interfaces"
)

type AuthorUseCase struct {
	repoAuthor interfaces.AuthorInterfaceUseCase
}

func NewUseCase(repoAuthor interfaces.AuthorInterfaceUseCase) *AuthorUseCase {
	return &AuthorUseCase{
		repoAuthor,
	}
}

func (uc *AuthorUseCase) GetAllAuthors(ctx context.Context, payload *author_proto.GetAllAuthorsRequest) (*author_proto.GetAllAuthorsResponse, error) {
	result, err := uc.repoAuthor.GetAllAuthors(ctx, payload)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (uc *AuthorUseCase) CreateAuthor(ctx context.Context, payload *author_proto.CreateAuthorRequest) (*author_proto.Author, error) {

	result, err := uc.repoAuthor.CreateAuthor(ctx, payload)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil

}

func (uc *AuthorUseCase) GetAuthorById(ctx context.Context, payload *author_proto.GetDetailAuthorRequest) (*author_proto.GetAuthorResponse, error) {

	result, err := uc.repoAuthor.GetAuthorById(ctx, payload)

	if err != nil {
		return nil, err
	}

	if result == nil {
		return &author_proto.GetAuthorResponse{}, lib.NewErrNotFound("author not found")
	}

	return result, nil
}

func (uc *AuthorUseCase) UpdateAuthor(ctx context.Context, payload *author_proto.UpdateAuthorRequest) (*author_proto.Author, error) {

	result, err := uc.repoAuthor.UpdateAuthor(ctx, payload)

	if err != nil {
		return nil, err
	}

	return result, nil

}

func (uc *AuthorUseCase) DeleteAuthor(ctx context.Context, id *author_proto.DeleteAuthorRequest) (*author_proto.DeleteAuthorResponse, error) {

	result, err := uc.repoAuthor.DeleteAuthor(ctx, id)

	if err != nil {
		return nil, err
	}

	return result, nil

}
