package main

import (
	//导入自动生成的proto包，并改名为pb
	pb "github.com/howl-io/shippy/consignment-service/proto/consignment"
	vesselPb "github.com/howl-io/shippy/vessel-service/proto/vessel"
	userPb "github.com/howl-io/shippy/user-service/proto/user"
	"log"
	"os"
	// for grpc
	//"google.golang.org/grpc"
	// for go-micro
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/client"
	"context"
	"errors"

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
		micro.WrapHandler(AuthWrapper),
	)

	server.Init()

	vesselCli := vesselPb.NewVesselServiceClient("go.micro.srv.vessel", server.Client())
	pb.RegisterShippingServiceHandler(server.Server(), &handler{session, vesselCli})

	if err := server.Run(); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}

// AuthWrapper 是一个高阶函数，入参是 "下一步" 函数，出参是认证函数
// 在返回的函数内部处理完认证逻辑后，再手动调用 fn() 进行下一步处理
// token 是从 consignment-cli 上下文取出的，再调用 user-service 将其做验证
// 认证通过则 fn() 继续执行，否则报错
func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in reques")
		}

		// Note this is now uppercase (not entirely sure why this is...)
		log.Printf("metadata: %v", meta)
		token := meta["Token"]

		// Auth here
		authClient := userPb.NewUserServiceClient("go.micro.srv.user", client.DefaultClient)
		authResp, err := authClient.ValidateToken(context.Background(), &userPb.Token {
			Token: token,
		})

		log.Println("Auth resp:", authResp)
		if err != nil {
			return err
		}

		err = fn(ctx, req, resp)
		return err
	}
}

