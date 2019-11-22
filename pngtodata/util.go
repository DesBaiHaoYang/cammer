package main

import (
	"encoding/base64"
	"io/ioutil"
	"os"

	"gopkg.in/mgo.v2/bson"
)

const UPLOAD_PATH string = "C:/Users/benben/Desktop/"

type Img struct {
	Id     bson.ObjectId `bson:"_id"`
	ImgUrl string        `bson:"imgUrl"`
}

func main() {
	cc, _ := ioutil.ReadFile("b.txt")
	dist, _ := base64.StdEncoding.DecodeString(string(cc))
	//写入新文件
	f, _ := os.OpenFile("b.png", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer f.Close()
	f.Write(dist)
}
