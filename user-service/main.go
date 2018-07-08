package main

import "log"
import (
	pb "github.com/heriyantoliu/Microservice/user-service/proto/user"
	"github.com/micro/go-micro"
	"fmt"
)

func main() {
	db, err := CreateConnection()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}
	defer db.Close()

	db.AutoMigrate(&pb.User{})
	repo := &UserRepository{db}
	TokenService := &TokenService{repo}

	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	srv.Init()

	pb.RegisterUserServiceHandler(srv.Server(), &service{repo, TokenService})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
