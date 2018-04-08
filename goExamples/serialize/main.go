package main

import (
	"fmt"
	"log"
	"myGoPractices/goExamples/serialize/pb"

	"github.com/golang/protobuf/proto"
)

func main() {
	p1 := pb.Person{
		Id:   1,
		Name: "hawken",
	}
	p2 := pb.Person{
		Id:   2,
		Name: "ssss",
	}
	p3 := pb.Person{
		Id:   3,
		Name: "zzz",
	}
	allPer := pb.AllPerson{
		Per: []*pb.Person{&p1, &p2, &p3},
	}
	// 对数据进行序列化
	data, err := proto.Marshal(&allPer)
	if err != nil {
		log.Fatalln("Marshal data error:", err)
	}
	fmt.Println(data)
	// 对已经序列化的数据进行反序列化
	var target pb.AllPerson
	err = proto.Unmarshal(data, &target)
	if err != nil {
		log.Fatalln("UnMarshal data error:", err)
	}
	fmt.Println(target.Per[1].GetName())
}
