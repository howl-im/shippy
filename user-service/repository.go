package main

import (
	pb "shippy/user-service/proto/user"
	"github.com/jinzhu/gorm"
	"log"
)



type IRepository interface {
	Get(id string) (*pb.User, error)
	GetAll()([]*pb.User, error)
	Create(*pb.User) error
	GetByEmailAndPasswrod(*pb.User) (*pb.User, error)
}

type Repository struct {
	db *gorm.DB
}

// 实现接口
func (repo *Repository) Get(id string) (*pb.User, error) {
	var user *pb.User
	user.Id = id
	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *Repository) GetAll() ([]*pb.User, error) {
	var users []*pb.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (repo *Repository) Create(user *pb.User) error {
	log.Printf("Create user: %v\n", user)
	if err := repo.db.Create(user).Error; err != nil {
		log.Printf("repo create user failed:%v\n", err)
		return err
	}

	return nil
}



func (repo *Repository) GetByEmailAndPasswrod(user *pb.User) (*pb.User, error) {
	if err := repo.db.Find(&user); err != nil {
		return nil, err.Error
	}

	return user, nil
}

