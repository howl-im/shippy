package main

import (
	pb "shippy/user-service/proto/user"
	"github.com/micro/go-micro"
	"log"
)


func main() {
	db, err := CreateConnection()
	defer db.Close()
	if err != nil {
		log.Fatalf("create connection failed: %v", err)
	}


	repo := &Repository{db}

	db.AutoMigrate(&pb.User{})

	server := micro.NewService(
					micro.Name("go.micro.srv.user"),
					micro.Version("latest"),
				)

	server.Init()

	pb.RegisterUserServiceHandler(server.Server(), &handler{repo})

	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
