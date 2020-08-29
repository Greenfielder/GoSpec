package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	pb "github.com/GreenFielder/GoSpec/lec5/client/proto/consignment"

	"google.golang.org/grpc"
)

const (
	address         = "localhost:50051"
	defaultFileName = "test1.json"
)

func parseJSON(file string) (*pb.Coefficients, error) {
	var coefficients *pb.Coefficients
	fileBody, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(fileBody, &coefficients)
	return coefficients, err
}

func main() {
	connection, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect to port: %v", err)
	}
	defer connection.Close()

	client := pb.NewSolverClient(connection)

	coefficients, err := parseJSON(defaultFileName)
	if err != nil {
		log.Fatalf("can not parse .json file: %v", err)
	}

	sol, err := client.Solve(context.Background(), coefficients)
	if err != nil {
		log.Fatalf("can not get response: %v", err)
	}
	log.Println(sol.NRoots)

	getAll, err := client.GetAll(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("can not get response: %v", err)
	}
	for _, v := range getAll.Solutions {
		log.Println(v)
	}
}
