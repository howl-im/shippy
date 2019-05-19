package main

import (
	"log"
	"gopkg.in/mgo.v2"
)

// 创建与 MongoDB 交互的主会话
func CreateSession(host string) (*mgo.Session, error) {
	log.Printf("host: %v\n", host)
	s, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	}

	s.SetMode(mgo.Monotonic, true)
	return s, nil
}

