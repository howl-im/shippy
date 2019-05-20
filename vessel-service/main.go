package main

import (
	pb "github.com/howl-io/shippy/vessel-service/proto/vessel"
	"os"
	"github.com/micro/go-micro"
	"log"
)

const (
	DB_DEFAULT_HOST = "172.17.0.1:27017"
)


func main() {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = DB_DEFAULT_HOST	
	}

	session, err := CreateSession(dbHost)
	defer session.Close()
	if err != nil {
		log.Fatalf("create session failed: %v", err)
	}

	

	repo := &Repository{session.Copy()}
	CreateDummyVessels(repo)

	server := micro.NewService(
					micro.Name("go.micro.srv.vessel"),
					micro.Version("latest"),
				)

	server.Init()

	pb.RegisterVesselServiceHandler(server.Server(), &handler{session})

	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func CreateDummyVessels(repo IRepository) {
	defer repo.Close()
	vessels := []*pb.Vessel {
		{
			Id: "vessel003",
			Name: "Boaty McBoatface",
			MaxWeight: 200000,
			Capacity: 500,
		},
	}

	for _, v := range vessels {
		repo.Create(v)
	}
}

