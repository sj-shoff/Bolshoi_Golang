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
	s.Set("key0", "2006")
	s.Set("key1", "2024")
	s.Set("key2", "year")
	res0 := s.Get("key1")
	res1 := s.Get("key1")
	res2 := s.Get("key2")

	fmt.Println(s)
	fmt.Println(res0)
	fmt.Println(res1)
	fmt.Println(res2)

}

// дз: тесты сет и гет (подсказка( можно подавать странный тип ))
