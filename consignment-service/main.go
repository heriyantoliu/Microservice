package main

import (
	pb "github.com/heriyantoliu/Microservice/consignment-service/proto/consignment"
	"github.com/micro/go-micro"
	"fmt"
	vesselProto "github.com/heriyantoliu/Microservice/vessel-service/proto/vessel"
	"os"
	"log"
)

const defaultHost  = "localhost:27017"

func main() {
	host := os.Getenv("DB_HOST")
	if host == ""{
		host = defaultHost
	}

	session, err := CreateSession(host)
	if err != nil {
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}
	defer session.Close()

	srv := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	vesselClient := vesselProto.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())

	srv.Init()

	pb.RegisterShippingServiceHandler(srv.Server(), &handler{session, vesselClient})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
