package main

import (
	context "context"
	"../proto"
	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}


func (s *server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error){

	a, b := request.GetA(), request.GetB()
	result := a + b

	return &proto.Response{Result : result}, nil
}

func (s *server) Multiply(ctx context.Context, request *proto.Request) (*proto.Response, error){
	
	a, b := request.GetA(), request.GetB()
	result := a * b

	return &proto.Response{Result : result}, nil
}

func main(){

	listener, err := net.Listen("tcp", "localhost:3000")
	if err != nil{
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil{
		panic(err)
	} 

}