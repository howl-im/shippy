package main

import (
	pb "github.com/howl-io/shippy/user-service/proto/user"
	"log"
	"context"
	"golang.org/x/crypto/bcrypt"
	"errors"
	//_ "github.com/micro/go-micro/go-plugins/broker/nats"
	"github.com/micro/go-micro"

)

const (
	TOPIC = "user.created"
)


// 定义货轮微服务
type handler struct {
	repo IRepository
	tokenService IAuthable
	//PubSub broker.Broker
	Publisher micro.Publisher
}


// 实现服务端
func (h *handler) Create(ctx context.Context, req *pb.User, resp *pb.Response) error {
	log.Printf("handler req: %v\n", req)
	hashedPword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	req.Password = string(hashedPword)
	if err := h.repo.Create(req); err != nil {
		log.Printf("handler create user failed:%v\n", err)
		return err
	}

	log.Printf("after handler req: %v\n", req)
	resp.User = req


	// 发布带有用户所有信息的消息
//	if err := h.publishEvent(req); err != nil {
//		return err
//	}

	if err := h.Publisher.Publish(ctx, req); err != nil {
		return err
	}

	return nil
}

// 发送消息通知
//func (h *handler) publishEvent(user *pb.User) error {
//	body, err := json.Marshal(user)
//	if err != nil {
//		return err
//	}
//
//	msg := &broker.Message {
//		Header: map[string]string {
//			"id": user.Id,
//		},
//		Body: body,
//	}
//
//	if err := h.PubSub.Publish(TOPIC, msg); err != nil {
//		log.Fatalf("[pub] failed: %v\n", err
//	}
//
//	return nil
//}


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
	log.Printf("calling Auth\n")
	log.Printf("req: %v\n", req)
	user, err := h.repo.GetByEmail(req.Email)
	if err != nil {
		log.Printf("get user by email failed: %v", err)
		return err
	}

	log.Printf("user: %v\n", user)

	log.Printf("compare hash and password\n")
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		log.Printf("auth failed: %v", err)
		return err
	}

	log.Printf("encode \n")
	t, err := h.tokenService.Encode(user)
	if err != nil {
		log.Printf("encode failed: %v", err)
		return err
	}


	log.Printf("set resp token\n")
	resp.Token = t
	return nil
}

func (h *handler) ValidateToken(ctx context.Context, req *pb.Token, resp *pb.Token) error {
	log.Printf("calling ValidateToken, req.Taken: %v\n", req.Token)
	// Decode token
	claims, err := h.tokenService.Decode(req.Token)
	if err != nil {
		log.Printf("decode failed: %v\n", err)
		return err
	}

	if claims.User.Id == "" {
		log.Printf("user id is nil\n")
		return errors.New("invalid user")
	}

	resp.Valid = true
	return nil
}
