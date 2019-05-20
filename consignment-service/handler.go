package main

import (
	//导入自动生成的proto包，并改名为pb
	pb "github.com/howl-io/shippy/consignment-service/proto/consignment"
	vesselPb "github.com/howl-io/shippy/vessel-service/proto/vessel"
	"context"
	"log"
	// for grpc
	//"google.golang.org/grpc"
	// for go-micro
	"gopkg.in/mgo.v2"
)

//
// 定义微服务
//
type handler struct {
	session *mgo.Session
	vesselCli vesselPb.VesselServiceClient
}

// 从主会话中 Clone() 出新的会话处理查询
func (h *handler) GetRepo() IRepository {
	return &Repository{h.session.Clone()}
}


//
// handler 实现 consignment.pb.go 中的 ShippingServiceServer 接口
// 使 service 作为 gRPC 的服务端
//
// 托运新的货物
// for grpc
//func (h *handler) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
// for go-micro
func (h *handler) CreateConsignment(ctx context.Context, req *pb.Consignment, resp *pb.Response) (error) {
	// 接收承运的货物
	defer h.GetRepo().Close()
	log.Printf("CreateConsignment called by client\n")
	//log.Printf("req: %+v\n", req)

	spec := &vesselPb.Specification {
		Capacity: int32(len(req.Containers)),
		MaxWeight: req.Weight,
	}

	vesselResp, err := h.vesselCli.FindAvailable(context.Background(), spec)
	if err != nil {
		log.Printf("not found vessel to use: %v\n", err)
		return err
	}

	// 货物被承运
	log.Printf("found vessel: %s\n", vesselResp.Vessel.Name)
	req.VesselId = vesselResp.Vessel.Id


	//consignment, err := h.repo.Create(req)
	err = h.GetRepo().Create(req)
	if err != nil {
		return err
	}
	resp.Created = true
	resp.Consignment = req

	return nil
}

//
// 获取目前托运的所有货物
//
// for grpc
//func (h *handler) GetConsignments(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
// for micro
func (h *handler) GetConsignments(ctx context.Context, req *pb.GetRequest, resp *pb.Response) (error) {
	defer h.GetRepo().Close()
	consignments, err := h.GetRepo().GetAll()
	if err != nil {
		return err
	}

	resp.Consignments = consignments
	return nil
}

