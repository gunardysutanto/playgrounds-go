package main

import (
	"log"
	"net"
	pb "proto/teams.pb.go"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	listen, err := net.Listen("tcp","localhost:8090")
	if(er != nil) {
		log.Fatalf("Unable to listen the port: %v",err)
	}

	s := grpc.NewServer()
	pb.RegisterTeamServiceServer(s, &server{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	log.Printf("Server runs on: %s", listen.Addr().String())
}
