package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/GreenFielder/GoSpec/lec5/server/proto/consignment"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// Command = Coefficients, Response = Solution
//Solve = CreateCommand

type repository interface {
	GetSolve(*pb.Coefficients) (*pb.Coefficients, error)
	GetAllCoefficients() []*pb.Coefficients
}

//Repository ... Наша база данных
type Repository struct {
	coefficientes []*pb.Coefficients
}

//GetSolve ..
func (r *Repository) GetSolve(coefficients *pb.Coefficients) (*pb.Coefficients, error) {
	updCoeffs := append(r.coefficientes, coefficients)
	r.coefficientes = updCoeffs
	return coefficients, nil
}

//GetAllCoefficients func
func (r *Repository) GetAllCoefficients() []*pb.Coefficients {
	return r.coefficientes
}

type service struct {
	repo repository
}

//Sqr func ...
func Sqr(a int32, b int32, c int32) (n int32) {
	var d int32

	if a >= 1 {
		d = (b * b) - (4 * a * c)
		if d > 0 {
			n = 2
		} else if d == 0 {
			n = 1
		} else {
			n = 0
		}
	} else {
		if b == 0 {
			n = 0
		} else {
			n = 1
		}
	}
	return n
}

func (s *service) Solve(ctx context.Context, req *pb.Coefficients) (*pb.Solution, error) {
	coefficients, err := s.repo.GetSolve(req)
	if err != nil {
		return nil, err
	}
	a := coefficients.A
	b := coefficients.B
	c := coefficients.C
	n := Sqr(a, b, c)
	return &pb.Solution{Coefs: coefficients, NRoots: n}, nil
}

func (s *service) GetAll(ctx context.Context, req *pb.GetRequest) (*pb.Solutions, error) {
	coefficientes := s.repo.GetAllCoefficients()
	var sol []*pb.Solution
	for _, v := range coefficientes {
		n := Sqr(v.A, v.B, v.C)
		//пока так не могу понять куда двигаться дальше
		fmt.Println(n)
	}

	return &pb.Solutions{Solutions: sol}, nil
}

func main() {
	repo := &Repository{}
	//Настройка gRPC сервера
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen port: %v", err)
	}
	server := grpc.NewServer()
	//Регистрируем наш сервис для сервера
	ourService := &service{repo}
	pb.RegisterSolverServer(server, ourService)
	//Чтобы выходные параметры сервера сохранялись в go-runtime
	reflection.Register(server)
	log.Println("gRPC server running on port:", port)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to server from port: %v", err)
	}
}
