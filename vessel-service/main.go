package main

import (
	pb "github.com/heriyantoliu/Microservice/vessel-service/proto/vessel"
	"github.com/micro/go-micro"
	"fmt"
	"os"
	"log"
)

const defaultHost  = "localhost:27017"

func createDummyData(repo Repository) {
	defer repo.Close()
	vessels := []*pb.Vessel{
		{
			Id:"vessel001",
			Name:"Kane's Salty Secret",
			MaxWeight:200000,
			Capacity:500,
		},
	}
	for _, v := range vessels {
		repo.Create(v)
	}
}

func main() {

	host := os.Getenv("DB_HOST")

	if host == ""{
		host = defaultHost
	}

	session, err := CreateSession(host)
	if err != nil {
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}

	srv := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)

	repo := &VesselRepository{session.Copy()}
	createDummyData(repo)

	srv.Init()

	pb.RegisterVesselServiceHandler(srv.Server(), &service{session})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
