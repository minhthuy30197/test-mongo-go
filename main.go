package main

import (
	"log"
	"time"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	hosts      = "localhost:27017"
	database   = "test_db"
	username   = "user1"
	password   = "example"
	collection = "post"
)

type Post struct {
	Tile    string
	Content string
}

func main() {
	// Tạo phiên kết nối với MongDB
	info := &mgo.DialInfo{
		Addrs:    []string{hosts},
		Timeout:  60 * time.Second,
		Database: database,
		Username: username,
		Password: password,
	}
	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	}

	// Tạo collection
	col := session.DB(database).C(collection)

	// Ghi dữ liệu
	err = col.Insert(&Post{"Hello World", "Đây là bài viết hướng dẫn thao tác với MongoDB trong Golang"},
		&Post{"Thời tiết", "Hôm nay nắng đẹp quá!!!"})
	if err != nil {
		log.Fatal(err)
	}

	// Đọc dữ liệu
	var posts []Post
	err = col.Find(bson.M{}).All(&posts)
	if err != nil {
		log.Fatal(err)
	}
	for _, post := range posts {
		log.Println(post)
	}
}
