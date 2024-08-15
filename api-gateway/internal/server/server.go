package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sharing_vasion_indonesia/api_gateway/internal/article/delivery"
	"sharing_vasion_indonesia/api_gateway/internal/article/usecase"
	article_proto "sharing_vasion_indonesia/pkg/proto"
	"syscall"
	"time"

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

	articleServicePort := ":9001"
	articleServiceconn, err := grpc.DialContext(ctx, articleServicePort, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer articleServiceconn.Close()

	articleClient := article_proto.NewArticleServiceClient(articleServiceconn)

	usecase := usecase.NewArticleUseCase(articleClient)
	delivery := delivery.NewDelivery(usecase, s.mux)
	delivery.Routes()

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
