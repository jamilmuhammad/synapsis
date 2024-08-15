package main

import (
	"context"
	"log"
	"net"
	"user-service/database"
	"user-service/internal/delivery"
	"user-service/internal/repository"
	"user-service/internal/usecase"
	"user-service/userpb"

	"github.com/jackc/pgx/v4/pgxpool"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func setupPostgres(dbURL string) {
	// Create user and database
	// createUserCmd := exec.Command("createuser", "-s", os.Getenv("DB_USER"))
	// createUserCmd.Run()

	// createDbCmd := exec.Command("createdb", "-O", os.Getenv("DB_USER"), os.Getenv("DB_NAME"))
	// createDbCmd.Run()

	// Connect to the database and create the users table
	// dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_PORT"),
	// 	os.Getenv("DB_NAME"),
	// )

	conn, err := pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close()

	_, err = conn.Exec(context.Background(), `
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            username VARCHAR(50) UNIQUE NOT NULL,
            email VARCHAR(100) UNIQUE NOT NULL,
            password VARCHAR(255) NOT NULL,
            created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
        )
    `)
	if err != nil {
		log.Fatalf("Failed to create users table: %v", err)
	}

	log.Println("PostgreSQL setup completed")
}

func main() {
	// Database connection
	//db, err := sql.Open("postgres", "postgres://username:password@localhost/userdb?sslmode=disable")
	//if err != nil {
	//	log.Fatalf("Failed to connect to database: %v", err)
	//}
	//defer db.Close()

	// const projectDirName = "user-service" // change to relevant project name

	// projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	// currentWorkDirectory, _ := os.Getwd()
	// rootPath := projectName.Find([]byte(currentWorkDirectory))

	// err := godotenv.Load("../.env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Error loading .env file: %v", err)
	// }
	// cfg := lib.LoadConfigByFile("./cmd", "config", "yaml")

	// // Construct database connection string
	// dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
	// 	cfg.DB.DB_USER,
	// 	cfg.DB.DB_DB_PASSWORD,
	// 	cfg.DB.DB_DB_HOST,
	// 	cfg.DB.DB_DB_PORT,
	// 	cfg.DB.DB_DB_NAME,
	// )
	// const DATABASE_URL string = "postgres://user:pass@userdb:5432/userdb"

	// log.Printf(dbURL)

	// setupPostgres(DATABASE_URL)

	// // Create a connection pool
	// config, err := pgxpool.ParseConfig(DATABASE_URL)
	// if err != nil {
	// 	log.Fatalf("Unable to parse connection string: %v\n", err)
	// }

	// conn, err := pgxpool.ConnectConfig(context.Background(), config)
	// if err != nil {
	// 	log.Fatalf("Unable to connect to database: %v\n", err)
	// }
	// defer conn.Close()

	db, err := database.InitDatabase()

	if err != nil {
		log.Fatal(err)
	}

	// Create repository, usecase, and handler
	userRepo := repository.NewUserRepository(db)
	userService := usecase.NewUseCase(userRepo)
	userDelivery := delivery.NewDelivery(userService)
	// userHandler := handler.NewUserHandler(userService)

	// Create gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, userDelivery)

	log.Println("Starting gRPC server on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
