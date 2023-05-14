package main

import (
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/tumbleweedd/mediasoft-intership/customer-service/pkg/rabbitmq"
	"github.com/tumbleweedd/mediasoft-intership/customer-service/pkg/repository"
	"github.com/tumbleweedd/mediasoft-intership/customer-service/pkg/service"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error getting env, %v", err)
	}

	db, err := repository.NewPostgresDB(&repository.Config{
		PgHost:         os.Getenv("DB_HOST"),
		PgPort:         os.Getenv("DB_PORT"),
		PgUser:         os.Getenv("DB_USER"),
		PgPwd:          os.Getenv("DB_PASSWORD"),
		PgDBName:       os.Getenv("DB_NAME"),
		PgDBSchemaName: os.Getenv("DB_SCHEMA_NAME"),
		PgSSLMode:      os.Getenv("DB_SSLMODE"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	lis, err := net.Listen("tcp", os.Getenv("PORT"))
	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Auth Svc on", os.Getenv("PORT"))

	s := grpc.NewServer()

	rabbitmqConn, err := rabbitmq.NewRabbitMQConn("guest", "guest", "localhost", "5672", "queue_name")
	if err != nil {
		log.Fatalln("Failed to rabbitMQ conn: ", err)
	}
	defer rabbitmqConn.Close()

	repo := repository.NewRepository(db)
	svc := service.NewService(repo, rabbitmqConn)

	customer.RegisterOfficeServiceServer(s, svc)
	customer.RegisterUserServiceServer(s, svc)
	customer.RegisterOrderServiceServer(s, svc)

	if err := s.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
