package main

import (
	"fmt"
	"log"

	go_protobuf_manual "github.com/yuttasakcom/go-protobuf-manual/github.com/yuttasakcom/go-protobuf-manual"
	"google.golang.org/protobuf/proto"
)

func main() {
	book := &go_protobuf_manual.Book{
		Isbn: 123456,
		Name: "Go Programming",
	}

	data, err := proto.Marshal(book)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	newBook := &go_protobuf_manual.Book{}
	err = proto.Unmarshal(data, newBook)

	if err != nil {
		log.Fatal("Unmarshaling error: ", err)
	}

	fmt.Println("Book name: ", newBook.GetName())
	fmt.Println("ISBN: ", newBook.GetIsbn())
}
