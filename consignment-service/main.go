package main

import (
	//导入自动生成的proto包，并改名为pb
	pb "shippy/consignment-service/proto/consignment"
	vesselPb "shippy/vessel-service/proto/vessel"
	"log"
	"os"
	// for grpc
	//"google.golang.org/grpc"
	// for go-micro
	"github.com/micro/go-micro"
)

const (
	DEFAULT_MDB_HOST = "172.17.0.1:27017"
)


func main() {
//	listener, err := net.Listen("tcp", PORT)
//	if err != nil {
//		log.Fatalf("failed to listen: %v", err)
//	}
//
//	log.Printf("listen on port:%s\n", PORT)
//
//	server := grpc.NewServer()
//	repo := Repository{}
//
//	// 向 gRPC 注册微服务
//	// 此时会把已经实现的微服务 service 与协议中的 ShippingService 绑定
//	pb.RegisterShippingServiceServer(server, &service{repo})
//
//	if err := server.Serve(listener); err != nil {
//		log.Fatalf("failed to serve: %v", err)
//	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = DEFAULT_MDB_HOST
	}

	session, err := CreateSession(dbHost)
	if err != nil {
		log.Fatalf("create session error: %v\n", err)
	}

	defer session.Close()
	
	server := micro.NewService ( 
		// 必须与 consignment.proto 中的包一致
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	server.Init()

	vesselCli := vesselPb.NewVesselServiceClient("go.micro.srv.vessel", server.Client())
	pb.RegisterShippingServiceHandler(server.Server(), &handler{session, vesselCli})

	if err := server.Run(); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}



