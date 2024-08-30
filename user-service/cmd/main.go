package main

import (
	"context"
	"fmt"
	"lib"
	"syscall"
	"time"
	"log"
	"net"
	"os"
	"os/signal"
	"path/filepath"

	"user-service/internal/handler"
	"user-service/internal/repository"
	"user-service/internal/usecase"
	"user-service/userpb"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	port := ":9001"

	cfg := lib.LoadConfigByFile("/app/user-service/cmd/", "config", "yml")

	// Construct database connection string
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.DB.DB_USER,
		cfg.DB.DB_PASSWORD,
		cfg.DB.DB_HOST,
		cfg.DB.DB_PORT,
		cfg.DB.DB_NAME,
	)

	// // Create a connection pool
	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		log.Fatalf("Unable to parse connection string: %v\n", err)
	}

	conn, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	// defer conn.Close()

	// conn, err := database.InitDatabase(context.Background(), dbURL)

	// if err != nil {
	// 	log.Fatal(err)
	// }

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
	userRepo := repository.NewUserRepository(conn)
	userService := usecase.NewUseCase(userRepo)
	userHandler := handler.NewHandler(userService)
	// userHandler := handler.NewUserHandler(userService)

	s := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: 5 * time.Minute,
		Timeout:           15 * time.Second,
		MaxConnectionAge:  5 * time.Minute,
		Time:              5 * time.Minute,
	}))
	user.RegisterUserServiceServer(s, userHandler)

	go func() {
		err := s.Serve(l)
		if err != nil {
			log.Println(err)
			cancel()
		}
	}()

	log.Printf("GRPC Server User is listening on port: %v", port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)

	select {
	case q := <-quit:
		log.Println("signal.Notify:", q)
	case done := <-ctx.Done():
		log.Println("ctx.Done:", done)
	}

	log.Println("Server GRPC User Exited Properly")
}
