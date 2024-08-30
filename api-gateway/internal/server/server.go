package server

import (
	"api-gateway/internal/user/handler"
	"api-gateway/internal/user/usecase"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"user-service/userpb"

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

	userUsecase := usecase.NewUserUseCase(userClient)
	userHandler := handler.NewUserHandler(userUsecase, s.mux)
	userHandler.Routes()

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
