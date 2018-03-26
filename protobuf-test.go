package main

/*
simple data marshal-unmarshal test
*/
import (
	"fmt"
	"log"
	"github.com/golang/protobuf/proto"
	buf "go-protobuf-test/message"
)

func main() {
	fmt.Println("protobuf test")
	// declare a new protobuf message
	// this type is automatically generated from .proto file by protoc
	p := &buf.Person{ 
		Name: "Gnu Shim",
		Id:1,
		Email: "geunwoo.j.shim@gmail.com",
	}
	fmt.Println("Original message :\n\t",p)
	out, marshalErr := proto.Marshal(p)
	// marshal returns bytes array
	fmt.Println("marshal bytes :\n\t", out, marshalErr)
	if marshalErr != nil{
		log.Fatal("error occured during marshaling : ", marshalErr)
	}
	// declare a new blank protobuf message
	p2 := &buf.Person{}
	// unmarshal bytes array to protobuf message and assign it
	unmarshalErr := proto.Unmarshal(out, p2)
	if unmarshalErr != nil{
		log.Fatal("error occured during unmarshaling : ", unmarshalErr)
	}
	// unmarshal message
	fmt.Println("unmarshal message :\n\t", p2)
	
}
