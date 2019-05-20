package main

import (
	//导入自动生成的proto包，并改名为pb
	pb "github.com/howl-io/shippy/consignment-service/proto/consignment"
	"log"
	"gopkg.in/mgo.v2"
)

const (
	PORT            = ":50051"
	DB_NAME         = "shippy"
	CON_COLLECTION  = "consignments"
)


//
// 仓库接口
//
type IRepository interface {
	Create(consignment *pb.Consignment) (error)
	GetAll() ([]*pb.Consignment, error)
	Close() 
}

//
// 存放多批货物的仓库，实现了IRepository接口
//
type Repository struct {
	session *mgo.Session
}

func (repo *Repository) Create(consignment *pb.Consignment) (error) {
	log.Printf("Create called by client\n")
	//log.Printf("consignment: %+v\n", consignment)
	err := repo.collection().Insert(consignment)
	if err != nil {

		log.Printf("create consignment failed: %v\n", err)
	}

	return err

}

func (repo *Repository) GetAll() ([]*pb.Consignment, error) {
	var consignments []*pb.Consignment
	// Find() 一般用来执行查询，如果想执行 select * 则直接传入 nil 即可
	// 通过 .All() 将查询结果绑定到 consignments 变量上
	// 对应的 .One() 则只读取一行记录
	err := repo.collection().Find(nil).All(&consignments)

	return consignments, err
}

// 关闭连接
func (repo *Repository) Close() {
	// Close() 会在每次查询结束的时候关闭会话
	// Mgo 会在启动的时候生成一个 "主" 会话
	// 你可以使用 Copy() 直接从主会话复制出新的会话来执行，即每个查询都会有自己的数据库会话
	// 同时每个会话都有自己连接到数据库的 socket 及错误处理，这么做既安全又高效
	// 如果只使用一个连接到数据库的主 socket 来执行查询，那很多请处理都会阻塞
	// Mgo 因此能在不使用锁的情况下完美处理并发请求
	// 不过弊端就是，每次查询结束之后，必须确保数据库会话要收到 Close
	// 否则将建立过多无用的连接，白白浪费数据库资源
	repo.session.Close()
}

// 获取数据库连接
func (repo *Repository) collection() *mgo.Collection {
	return repo.session.DB(DB_NAME).C(CON_COLLECTION)
}

