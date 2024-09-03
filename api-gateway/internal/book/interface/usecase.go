package interfaces

import (
	book_proto "book-service/bookpb/book"
	"context"
)

type BookInterfaceUseCase interface {
	GetAllBooks(ctx context.Context, payload *book_proto.GetAllBooksRequest) (*book_proto.GetAllBooksResponse, error)
	CreateBook(ctx context.Context, post *book_proto.CreateBookRequest) (*book_proto.Book, error)
	UpdateBook(ctx context.Context, post *book_proto.UpdateBookRequest) (*book_proto.Book, error)
	GetBookById(ctx context.Context, payload *book_proto.GetDetailBookRequest) (*book_proto.GetBookResponse, error)
	DeleteBook(ctx context.Context, id *book_proto.DeleteBookRequest) (*book_proto.DeleteBookResponse, error)
}
