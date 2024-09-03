package main

import (
	"context"
	"fmt"
	"lib"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"book-service/bookpb/book"
	"book-service/database"
	"book-service/internal/handler"
	"book-service/internal/repository"
	"book-service/internal/usecase"

	author "author-service/authorpb"
	category "category-service/categorypb"

	_ "github.com/lib/pq"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	port := ":9004"

	cfg := lib.LoadConfigByFile("/app/book-service/cmd", "config", "yml")

	// Construct database connection string
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.DB.DB_USER,
		cfg.DB.DB_PASSWORD,
		cfg.DB.DB_HOST,
		cfg.DB.DB_PORT,
		cfg.DB.DB_NAME,
	)

	log.Println(dbURL)

	// defer conn.Close()

	conn, err := database.InitDatabase(ctx, dbURL)

	if err != nil {
		log.Fatal(err)
	}

	// Create gRPC server
	l, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	defer l.Close()

	tracer, closer, err := lib.InitJaeger(cfg.App.Name)
	if err != nil {
		log.Fatalf("could not initialize jaeger tracer: %v", err)
	}
	defer closer.Close()

	opentracing.SetGlobalTracer(tracer)

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

	// Create repository, usecase, and handler
	bookRepo := repository.NewBookRepository(conn, authorClient, categoryClient)
	bookService := usecase.NewUseCase(bookRepo)
	bookHandler := handler.NewHandler(bookService)

	s := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: 5 * time.Minute,
		Timeout:           15 * time.Second,
		MaxConnectionAge:  5 * time.Minute,
		Time:              5 * time.Minute,
	}))
	book.RegisterBookServiceServer(s, bookHandler)

	go func() {
		err := s.Serve(l)
		if err != nil {
			log.Println(err)
			cancel()
		}
	}()

	log.Printf("GRPC Server Book is listening on port: %v", port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)

	select {
	case q := <-quit:
		log.Println("signal.Notify:", q)
	case done := <-ctx.Done():
		log.Println("ctx.Done:", done)
	}

	log.Println("Server GRPC Book Exited Properly")
}
