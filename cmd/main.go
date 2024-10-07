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

	s.Set("key1", "2024")
	s.Set("key2", "year")
	res1 := s.Get("key1")
	kind1 := s.GetKind("key1")
	res2 := s.Get("key2")
	kind2 := s.GetKind("key2")
	fmt.Println(s)

	fmt.Println(*res1, kind1)
	fmt.Println(*res2, kind2)
}
