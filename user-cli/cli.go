package main
import (
	pb "github.com/howl-io/shippy/user-service/proto/user"
	_ "github.com/micro/go-micro"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	_ "github.com/micro/cli"
	"log"
	"context"
)



func main() {
//	// 连接到 gRPC 服务器
//	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure())
//	if err != nil {
//		log.Fatalf("connect error: %v", err)
//	}
//
//	defer conn.Close()
//
//
//	// 初始化 gRPC 客户端
//	client := pb.NewShippingServiceClient(conn)
	cmd.Init()

	client := pb.NewUserServiceClient("go.micro.srv.user", microclient.DefaultClient)
	
	name := "Ewan Valentime"
	email := "ewan.valentine89@gmail.com"
	passwrod := "test123"
	company := "BBC"

	resp, err := client.Create(context.TODO(), &pb.User {
		Name: name,
		Email: email,
		Password: passwrod,
		Company: company,
	})

	if err != nil {
		log.Fatalf("call Create error: %v", err)
	}

	log.Println("Created: ", resp.User.Id)

	resp, err = client.GetAll(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("could not list users: %v\n", err)
	}

	for _, v := range resp.Users {
		log.Println(v)
	}

	authResp, err := client.Auth(context.TODO(), &pb.User {
		Email: email,
		Password: passwrod,
	})

	if err != nil {
		log.Fatalf("auth failed: %v", err)
	}

	log.Println("token: ", authResp.Token)



//	// 创建 user_server 微服务的客户端
//	service := micro.NewService(
//		micro.Flags(
//			cli.StringFlag {
//				Name: "name",
//				Usage: "You full name",
//			},
//			cli.StringFlag {
//				Name: "email",
//				Usage: "Your email",
//			},
//			cli.StringFlag {
//				Name: "password",
//				Usage: "Your password",
//			},
//			cli.StringFlag {
//				Name: "company",
//				Usage: "Your company",
//			},
//		),
//	)
//
//	service.Init(
//		micro.Action(func(c *cli.Context) {
//			name := c.String("name")
//			email := c.String("email")
//			password := c.String("password")
//			company := c.String("company")
//
//			log.Printf("Name = %s, Email = %s, Password = %s, Company = %s\n", 
//							name, email, password, company)
//			r, err := client.Create(context.TODO(), &pb.User{
//				/*
//						Name: name,
//						Email: email,
//						Password: password,
//						Company: company,
//				*/
//						Name: "Lily",
//						Email: "lily@bbc.com",
//						Password: "testing123",
//						Company: "BBC",
//					})
//
//			if err != nil {
//				log.Fatalf("could not create: %v\n", err)
//			}
//
//			log.Printf("Create: %v\n", r.User.Id)
//
//			resp, err := client.GetAll(context.Background(), &pb.Request{})
//			if err != nil {
//				log.Fatalf("could not list users: %v\n", err)
//			}
//
//			for _, v := range resp.Users {
//				log.Println(v)
//			}
//
//			os.Exit(0)
//		}),
//	)
//
//	if err := service.Run(); err != nil {
//		log.Println(err)
//	}
	
	log.Printf("client exit...\n")
}


