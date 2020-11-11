package main

import (
	"fmt"

	"github.com/protocol-buffer/addressbook"
	"google.golang.org/protobuf/proto"
)

func main() {

	// encode
	book := &addressbook.AddressBook{}
	book.Person = []*addressbook.Person{
		&addressbook.Person{
			Name:  "andre",
			Id:    1,
			Email: "hokdre@gmail.com",
		},
	}
	fmt.Println(book)
	out, err := proto.Marshal(book)
	if err != nil {
		fmt.Printf("err marshaling : %s", err)
	}

	// decode
	decoded := &addressbook.AddressBook{}
	err = proto.Unmarshal(out, decoded)
	if err != nil {
		fmt.Printf("err unmarshaling : %s", err)
	}
	fmt.Println(decoded)

}
