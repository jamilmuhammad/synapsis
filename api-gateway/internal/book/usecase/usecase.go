package usecase

import (
	"context"

	book_proto "book-service/bookpb/book"
)

type bookUseCase struct {
	bookClient book_proto.BookServiceClient
}

func NewBookUseCase(bookClient book_proto.BookServiceClient) *bookUseCase {
	return &bookUseCase{bookClient}
}

func (u *bookUseCase) GetAllBooks(ctx context.Context, payload *book_proto.GetAllBooksRequest) (*book_proto.GetAllBooksResponse, error) {

	result, err := u.bookClient.GetAllBooks(ctx, payload)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *bookUseCase) CreateBook(ctx context.Context, payload *book_proto.CreateBookRequest) (*book_proto.Book, error) {

	result, err := u.bookClient.CreateBook(ctx, payload)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (uc *bookUseCase) GetBookById(ctx context.Context, payload *book_proto.GetDetailBookRequest) (*book_proto.GetBookResponse, error) {

	result, err := uc.bookClient.GetBookById(ctx, payload)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *bookUseCase) UpdateBook(ctx context.Context, payload *book_proto.UpdateBookRequest) (*book_proto.Book, error) {

	result, err := u.bookClient.UpdateBook(ctx, payload)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *bookUseCase) DeleteBook(ctx context.Context, payload *book_proto.DeleteBookRequest) (*book_proto.DeleteBookResponse, error) {

	result, err := u.bookClient.DeleteBook(ctx, payload)

	if err != nil {
		return nil, err
	}

	return result, nil
}
