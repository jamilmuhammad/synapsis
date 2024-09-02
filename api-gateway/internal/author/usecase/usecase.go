package usecase

import (
	"context"

	author_proto "author-service/authorpb"
)

type authorUseCase struct {
	authorClient author_proto.AuthorServiceClient
}

func NewAuthorUseCase(authorClient author_proto.AuthorServiceClient) *authorUseCase {
	return &authorUseCase{authorClient}
}

func (u *authorUseCase) GetAllAuthors(ctx context.Context, payload *author_proto.GetAllAuthorsRequest) (*author_proto.GetAllAuthorsResponse, error) {

	result, err := u.authorClient.GetAllAuthors(ctx, payload)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *authorUseCase) CreateAuthor(ctx context.Context, payload *author_proto.CreateAuthorRequest) (*author_proto.Author, error) {

	result, err := u.authorClient.CreateAuthor(ctx, payload)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (uc *authorUseCase) GetAuthorById(ctx context.Context, payload *author_proto.GetDetailAuthorRequest) (*author_proto.GetAuthorResponse, error) {

	result, err := uc.authorClient.GetAuthorById(ctx, payload)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *authorUseCase) UpdateAuthor(ctx context.Context, payload *author_proto.UpdateAuthorRequest) (*author_proto.Author, error) {

	result, err := u.authorClient.UpdateAuthor(ctx, payload)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *authorUseCase) DeleteAuthor(ctx context.Context, payload *author_proto.DeleteAuthorRequest) (*author_proto.DeleteAuthorResponse, error) {

	result, err := u.authorClient.DeleteAuthor(ctx, payload)

	if err != nil {
		return nil, err
	}

	return result, nil
}
