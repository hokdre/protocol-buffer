package main

import (
	"log"

	"github.com/protocol-buffer/pb"
	"google.golang.org/protobuf/proto"
)

func main() {
	//1. marshaling
	laptop := pb.Laptop{
		Id:    "123",
		Brand: "Toshiba",
		Weight: &pb.Laptop_WeightKg{
			WeightKg: 10.2,
		},
	}

	laptop.Weight = &pb.Laptop_WeightLb{
		WeightLb: 10.4,
	}
	binaryLaptop, err := proto.Marshal(&laptop)
	if err != nil {
		log.Fatalf("Failed marshaling protoc : %s", err)
		return
	}

	//2 unmarshaling
	unMarshalLaptop := pb.Laptop{}
	if err := proto.Unmarshal(binaryLaptop, &unMarshalLaptop); err != nil {
		log.Fatalf("Failed unmarshaling protoc : %s", err)
		return
	}
}
