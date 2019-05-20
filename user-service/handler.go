package main

import (
	pb "github.com/howl-io/shippy/user-service/proto/user"
	"log"
	"context"
)


// 定义货轮微服务
type handler struct {
	repo IRepository
}


// 实现服务端
func (h *handler) Create(ctx context.Context, req *pb.User, resp *pb.Response) error {
	log.Printf("handler req: %v\n", req)
	if err := h.repo.Create(req); err != nil {
		log.Printf("handler create user failed:%v\n", err)
		return err
	}

	log.Printf("after handler req: %v\n", req)
	resp.User = req
	return nil
}

func (h *handler) Get(ctx context.Context, req *pb.User, resp *pb.Response) error {
	user, err := h.repo.Get(req.Id)
	if err != nil {
		return err
	}

	resp.User = user

	return nil
}

func (h *handler) GetAll(ctx context.Context, req *pb.Request, resp *pb.Response) error {
	users, err := h.repo.GetAll()
	if err != nil {
		return err
	}

	resp.Users = users

	return nil
}

func (h *handler) Auth(ctx context.Context, req *pb.User, resp *pb.Token) error {
	_, err := h.repo.GetByEmailAndPasswrod(req)
	if err != nil {
		return err
	}

	resp.Token = "`x_2nam"
	return nil
}

func (h *handler) ValidateToken(ctx context.Context, req *pb.Token, resp *pb.Token) error {
	return nil
}
