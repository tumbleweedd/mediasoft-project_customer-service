package app

import (
	"fmt"
	"github.com/joho/godotenv"
	repository2 "github.com/tumbleweedd/mediasoft-intership/customer-service/internal/repository"
	"github.com/tumbleweedd/mediasoft-intership/customer-service/internal/service"
	"github.com/tumbleweedd/mediasoft-intership/customer-service/pkg/broker/kafka/producer"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

const broker = "192.168.0.109:9092"

func Run() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error getting env, %v", err)
	}

	db, err := repository2.NewPostgresDB(&repository2.Config{
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

	kafkaProducer, err := producer.NewProducer(broker)
	if err != nil {
		log.Fatalln("Failed to kafka conn: ", err)
	}
	defer kafkaProducer.Close()

	done := make(chan struct{})
	defer close(done)

	repo := repository2.NewRepository(db)
	svc := service.NewService(repo, kafkaProducer, done)

	customer.RegisterOfficeServiceServer(s, svc)
	customer.RegisterUserServiceServer(s, svc)
	customer.RegisterOrderServiceServer(s, svc)

	if err := s.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
