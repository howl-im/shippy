package main

import (
	pb "shippy/vessel-service/proto/vessel"
	"gopkg.in/mgo.v2"
	 "gopkg.in/mgo.v2/bson"
	"log"
)

const (
	DB_NAME = "shippy"
	VESSEL_COLLECTION = "vessels"
)

type IRepository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	Create(*pb.Vessel) error
	Close()
}

type Repository struct {
	session *mgo.Session
}

// 实现接口
func (repo *Repository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	// 选择最近一艘容量、载重都符合的货轮
	var v *pb.Vessel
	log.Printf("find FindAvailable \n")
	
	err := repo.collection().Find(bson.M {
		"capacity": bson.M {"$gte": spec.Capacity},
		"maxweight": bson.M {"$gte": spec.MaxWeight},
	}).One(&v)

	if err != nil {
		log.Printf("Find failed: %v\n", err)
		return nil, err
	}
/*
	var vs []*pb.Vessel
	err := repo.collection().Find(nil).All(&vs)

	if err != nil {
		return nil, err
	}

	log.Printf("spec: %v\n", spec)
	log.Printf("all vessels: %v\n", vs)

	for _, ve := range vs {
		if ve.Capacity >= spec.Capacity && ve.MaxWeight >= spec.MaxWeight {
			log.Printf("found one:%v\n", ve)
			v = ve
			break;
		}

	}
*/

	log.Printf("spec: %v\n", spec)

	return v, nil
}

func (repo *Repository) Create(v *pb.Vessel) error {
	log.Printf("Insert: %v\n", v)
	return repo.collection().Insert(v)
}

func (repo *Repository) Close() {
	repo.session.Close()
}

func (repo *Repository) collection() *mgo.Collection {
	return repo.session.DB(DB_NAME).C(VESSEL_COLLECTION)
}

