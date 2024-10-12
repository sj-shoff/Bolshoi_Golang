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

	fmt.Println(s)

}

// дз: тесты сет и гет (подсказка( можно подавать странный тип ))
