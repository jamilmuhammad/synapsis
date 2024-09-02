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

	"author-service/authorpb"
	"author-service/database"
	"author-service/internal/handler"
	"author-service/internal/repository"
	"author-service/internal/usecase"

	_ "github.com/lib/pq"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	port := ":9003"

	cfg := lib.LoadConfigByFile("/app/author-service/cmd", "config", "yml")

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

	// Create repository, usecase, and handler
	authorRepo := repository.NewAuthorRepository(conn)
	authorService := usecase.NewUseCase(authorRepo)
	authorHandler := handler.NewHandler(authorService)

	s := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: 5 * time.Minute,
		Timeout:           15 * time.Second,
		MaxConnectionAge:  5 * time.Minute,
		Time:              5 * time.Minute,
	}))
	author.RegisterAuthorServiceServer(s, authorHandler)

	go func() {
		err := s.Serve(l)
		if err != nil {
			log.Println(err)
			cancel()
		}
	}()

	log.Printf("GRPC Server Author is listening on port: %v", port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)

	select {
	case q := <-quit:
		log.Println("signal.Notify:", q)
	case done := <-ctx.Done():
		log.Println("ctx.Done:", done)
	}

	log.Println("Server GRPC Author Exited Properly")
}
