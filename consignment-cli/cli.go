package main
import (
	pb "github.com/howl-io/shippy/consignment-service/proto/consignment"
	"io/ioutil"
	"encoding/json"
	"errors"
	//"google.golang.org/grpc"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/metadata"

	"log"
	"os"
	"context"
)

const (
	ADDRESS            = "localhost:50051"
	DEFAULT_INFO_FILE  = "consignment.json"
)

// 读取 consignment.json 中的记录的货物
func parseFile(fileName string) (*pb.Consignment, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var consignment *pb.Consignment
	err = json.Unmarshal(data, &consignment)
	if err != nil {
		return nil, errors.New("consignment.json file content error")
	}

	//log.Printf("consignment: %+v\n", consignment)
	return consignment, nil
}

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

	client := pb.NewShippingServiceClient("go.micro.srv.consignment", microclient.DefaultClient)

	// 在命令行中指定新的货物信息 json 文件
	infoFile := DEFAULT_INFO_FILE
	
	if len(os.Args) < 1 {
		log.Fatalln("Not enough arguments, expecting file and token.")
	}

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7IklkIjoiZGI1MTIzZDQtYjg4Ny00MTFhLWE1ZDgtYWRlMmUwMTExNzc2IiwiTmFtZSI6IkV3YW4gVmFsZW50aW1lIiwiQ29tcGFueSI6IkJCQyIsIkVtYWlsIjoiZXdhbi52YWxlbnRpbmU4OUBnbWFpbC5jb20iLCJQYXNzd29yZCI6IiQyYSQxMCROc3FFQ28zbXprVmxsRGFENjMwVmN1dXV6OTZGenNKOUcxU1JTdzdvNHkvakRRWDRhdXZJVyJ9LCJleHAiOjE1NTkwMjcyMzMsImlzcyI6ImdvLm1pY3JvLnNydi51c2VyIn0.A16BCPByhk1qkUkSYDPJ1kBXmtV-G6ydJjs6M7ByaFE"
	//infoFile = os.Args[1]
	//token := os.Args[1]
	//log.Printf("Args[1]: %v\n", os.Args[1])
	//log.Printf("Args[2]: %v\n", os.Args[2])

	// 解析货物信息
	consignment, err := parseFile(infoFile)
	if err != nil {
		log.Fatalf("parae info file error: %v", err)
	}

	// 创建带有用户 token 的context
	// consignment-service 服务端将从中取出 token， 解密取出用户身份
	tokenContext := metadata.NewContext(context.Background(), map[string]string {
		"Token": token,
	})

	// 调用 RPC
	// 将货物存储到仓库中
	resp, err := client.CreateConsignment(tokenContext, consignment)
	if err != nil {
		log.Fatalf("create consignment error: %v", err)
	}

	// 新货物是否拖运成功
	log.Printf("created: %t\n", resp.Created)

	log.Printf("Response: %+v\n", resp)	

	// 列出目前所有托运的货物
	resp, err = client.GetConsignments(tokenContext, &pb.GetRequest{})
	if err != nil {
		log.Fatalf("failed to list consignments: %v", err)
	}

	for _, c := range resp.Consignments {
		log.Printf("%+v", c)
	}

	log.Printf("client exit...\n")
}


