package gproto

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

func StartGrpcServer()  {
	//product *model.Product
	grpcServer := grpc.NewServer()
	//reflection.Register(grpcServer)

	/*productserver := &ProducService{
		Products: product,
	}
	pb.RegisterProductServiceServer(grpcServer,productserver)*/
	
	listener, err := net.Listen("tcp", "0.0.0.0:4000")
	if err != nil {
		log.Fatal("error loading gproto server")
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("error loading gproto server")
	}

}
