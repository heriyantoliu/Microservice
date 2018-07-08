package main

import pb "github.com/heriyantoliu/Microservice/user-service/proto/user"
import (
	microclient "github.com/micro/go-micro/client"
	"log"
	"golang.org/x/net/context"
	"os"
	"github.com/micro/go-micro/cmd"
)

func main() {
	cmd.Init()

	client := pb.NewUserServiceClient("go.micro.srv.user", microclient.DefaultClient)

	name:= "Ewan Valentine"
	email:= "ewan.valentine89@gmail.com"
	password:="test123"
	company:="BBC"

	r, err := client.Create(context.TODO(), &pb.User{
		Name: name,
		Email:email,
		Password:password,
		Company:company,
	})

	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Fatalf("Created: %s", r.User.Id)

	getAll, err := client.GetAll(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("Could not list users: %v", err)
	}

	for _, v := range getAll.Users {
		log.Println(v)
	}

	authResponse, err := client.Auth(context.TODO(), &pb.User{
		Email:email,
		Password:password,
	})

	if err != nil {
		log.Fatalf("Could not authenticate user: %s error: %v\n", email, err)
	}

	log.Printf("YOUR access token is: %s \n", authResponse.Token)
	os.Exit(0)
}
