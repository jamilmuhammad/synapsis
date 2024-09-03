package handler

import (
	"context"
	"log"

	book_proto "book-service/bookpb/book"
	"book-service/internal/interfaces"
)

type BookHandler struct {
	book_proto.UnimplementedBookServiceServer
	repoBook interfaces.BookInterfaceHandler
}

func NewHandler(repoBook interfaces.BookInterfaceHandler) *BookHandler {
	return &BookHandler{repoBook: repoBook}
}

func (uc *BookHandler) GetAllBooks(ctx context.Context, payload *book_proto.GetAllBooksRequest) (*book_proto.GetAllBooksResponse, error) {
	result, err := uc.repoBook.GetAllBooks(ctx, payload)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (uc *BookHandler) CreateBook(ctx context.Context, payload *book_proto.CreateBookRequest) (*book_proto.Book, error) {

	result, err := uc.repoBook.CreateBook(ctx, payload)

	if err != nil {
		return nil, err
	}

	return result, nil

}

func (uc *BookHandler) GetBookById(ctx context.Context, payload *book_proto.GetDetailBookRequest) (*book_proto.GetBookResponse, error) {

	result, err := uc.repoBook.GetBookById(ctx, payload)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil

}

func (uc *BookHandler) UpdateBook(ctx context.Context, payload *book_proto.UpdateBookRequest) (*book_proto.Book, error) {

	result, err := uc.repoBook.UpdateBook(ctx, payload)

	if err != nil {
		return nil, err
	}

	return result, nil

}

func (uc *BookHandler) DeleteBook(ctx context.Context, id *book_proto.DeleteBookRequest) (*book_proto.DeleteBookResponse, error) {

	result, err := uc.repoBook.DeleteBook(ctx, id)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil

}
