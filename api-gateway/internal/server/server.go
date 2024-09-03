package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	user_handler "api-gateway/internal/user/handler"
	user_usecase "api-gateway/internal/user/usecase"
	"user-service/userpb"

	category_handler "api-gateway/internal/category/handler"
	category_usecase "api-gateway/internal/category/usecase"
	"category-service/categorypb"

	author_handler "api-gateway/internal/author/handler"
	author_usecase "api-gateway/internal/author/usecase"
	"author-service/authorpb"

	book_handler "api-gateway/internal/book/handler"
	book_usecase "api-gateway/internal/book/usecase"
	"book-service/bookpb/book"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

type server struct {
	mux *mux.Router
	v   validator.Validate
}

func NewServer(mux *mux.Router) *server {
	return &server{mux: mux, v: *validator.New()}
}

func (s *server) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//user-service
	userServiceAddress := "user-service:9001"

	userServiceconn, err := grpc.Dial(
		userServiceAddress,
		grpc.WithInsecure(),
	)

	if err != nil {
		log.Fatal(err)
	}
	defer userServiceconn.Close()

	userClient := user.NewUserServiceClient(userServiceconn)

	userUsecase := user_usecase.NewUserUseCase(userClient)
	userHandler := user_handler.NewUserHandler(userUsecase, s.mux)
	userHandler.Routes()

	//category-service
	categoryServiceAddress := "category-service:9002"

	categoryServiceconn, err := grpc.Dial(
		categoryServiceAddress,
		grpc.WithInsecure(),
	)

	if err != nil {
		log.Fatal(err)
	}
	defer categoryServiceconn.Close()

	categoryClient := category.NewCategoryServiceClient(categoryServiceconn)

	categoryUsecase := category_usecase.NewCategoryUseCase(categoryClient)
	categoryHandler := category_handler.NewCategoryHandler(categoryUsecase, s.mux)
	categoryHandler.Routes()

	//author-service
	authorServiceAddress := "author-service:9003"

	authorServiceconn, err := grpc.Dial(
		authorServiceAddress,
		grpc.WithInsecure(),
	)

	if err != nil {
		log.Fatal(err)
	}
	defer authorServiceconn.Close()

	authorClient := author.NewAuthorServiceClient(authorServiceconn)

	authorUsecase := author_usecase.NewAuthorUseCase(authorClient)
	authorHandler := author_handler.NewAuthorHandler(authorUsecase, s.mux)
	authorHandler.Routes()

	//book-service
	bookServiceAddress := "book-service:9004"

	bookServiceconn, err := grpc.Dial(
		bookServiceAddress,
		grpc.WithInsecure(),
	)

	if err != nil {
		log.Fatal(err)
	}
	defer bookServiceconn.Close()

	bookClient := book.NewBookServiceClient(bookServiceconn)

	bookUsecase := book_usecase.NewBookUseCase(bookClient)
	bookHandler := book_handler.NewBookHandler(bookUsecase, s.mux)
	bookHandler.Routes()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)

	port := ":9000"

	go func() {
		server := &http.Server{
			Addr:         port,
			Handler:      s.mux,
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second,
		}

		err := server.ListenAndServe()
		if err != nil {
			log.Println(err)
			cancel()
		}
	}()

	log.Println("API Gateway listen on port", port)

	select {
	case q := <-quit:
		log.Println("signal.Notify:", q)
	case done := <-ctx.Done():
		log.Println("ctx.Done:", done)
	}

	log.Println("Server API Gateway Exited Properly")

	return nil
}
