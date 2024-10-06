package main

import (
	"fmt"
	"hello_world/iternal/pkg/storage"
	"log"
)

func main() {
	s, err := storage.NewStorage()

	if err != nil {
		log.Fatal(err)
	}

	s.Set("key2", "222")

	res1 := s.Get("key2")
	kind := s.GetKind("key2")

	fmt.Println(*res1, kind)
}
