package main

import (
	pb "github.com/howl-io/shippy/user-service/proto/user"
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

	// 初始化命令行环境
	server.Init()

	// 获取 broker 实例
	//pubSub := server.Server().Options().Broker
	publisher := micro.NewPublisher(TOPIC, server.Client())

	token := &TokenService{repo}

	pb.RegisterUserServiceHandler(server.Server(), &handler{repo, token, publisher})

	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

