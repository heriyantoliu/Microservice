package main

import (
	pb "github.com/heriyantoliu/Microservice/consignment-service/proto/consignment"
	"io/ioutil"
	"encoding/json"
	microclient "github.com/micro/go-micro/client"
	"log"
	"os"
	"context"
	"fmt"
	"github.com/micro/go-micro/cmd"
)

const (
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}
func main() {
	cmd.Init()

	client := pb.NewShippingServiceClient("go.micro.srv.consignment", microclient.DefaultClient)

	file := defaultFilename
	if len (os.Args) >1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(),consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range getAll.Consignments {
		fmt.Println(v)
	}
}
