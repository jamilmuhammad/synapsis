package usecase

import (
	"context"
	"lib"
	"log"

	book_proto "book-service/bookpb/book"
	"book-service/internal/interfaces"
)

type BookUseCase struct {
	repoBook interfaces.BookInterfaceUseCase
}

func NewUseCase(repoBook interfaces.BookInterfaceUseCase) *BookUseCase {
	return &BookUseCase{
		repoBook,
	}
}

func (uc *BookUseCase) GetAllBooks(ctx context.Context, payload *book_proto.GetAllBooksRequest) (*book_proto.GetAllBooksResponse, error) {
	result, err := uc.repoBook.GetAllBooks(ctx, payload)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (uc *BookUseCase) CreateBook(ctx context.Context, payload *book_proto.CreateBookRequest) (*book_proto.Book, error) {

	result, err := uc.repoBook.CreateBook(ctx, payload)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil

}

func (uc *BookUseCase) GetBookById(ctx context.Context, payload *book_proto.GetDetailBookRequest) (*book_proto.GetBookResponse, error) {

	result, err := uc.repoBook.GetBookById(ctx, payload)

	if err != nil {
		return nil, err
	}

	if result == nil {
		return &book_proto.GetBookResponse{}, lib.NewErrNotFound("book not found")
	}

	return result, nil
}

func (uc *BookUseCase) UpdateBook(ctx context.Context, payload *book_proto.UpdateBookRequest) (*book_proto.Book, error) {

	result, err := uc.repoBook.UpdateBook(ctx, payload)

	if err != nil {
		return nil, err
	}

	return result, nil

}

func (uc *BookUseCase) DeleteBook(ctx context.Context, id *book_proto.DeleteBookRequest) (*book_proto.DeleteBookResponse, error) {

	result, err := uc.repoBook.DeleteBook(ctx, id)

	if err != nil {
		return nil, err
	}

	return result, nil

}
