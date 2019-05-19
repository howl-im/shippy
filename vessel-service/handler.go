package main

import (
	pb "shippy/vessel-service/proto/vessel"
	"context"
	"gopkg.in/mgo.v2"
	"log"
)


// 定义货轮微服务
type handler struct {
	session *mgo.Session
}

func (h *handler) GetRepo() IRepository {
	return &Repository{h.session.Clone()}
}


// 实现服务端
func (h *handler) FindAvailable(ctx context.Context, spec *pb.Specification, resp *pb.Response) error {
	defer h.GetRepo().Close()
	v, err := h.GetRepo().FindAvailable(spec)
	if err != nil {
		log.Printf("find available failed: %v\n", err)
		return err
	}

	resp.Vessel = v
	return nil
}

func (h *handler) Create(ctx context.Context, req *pb.Vessel, resp *pb.Response) error {
	defer h.GetRepo().Close()
	if err := h.GetRepo().Create(req); err != nil {
		return err
	}

	resp.Vessel = req
	resp.Created = true

	return nil
}

