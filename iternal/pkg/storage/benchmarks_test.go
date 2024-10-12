package storage

import (
	"testing"
)

func BenchmarkSet(b *testing.B) {
	for _, tCase := range cases_for_set {
		b.Run(tCase.key, func(bb *testing.B) {
			s, _ := NewStorage()

			bb.ResetTimer()

			for i := 0; i < bb.N; i++ {
				s.Set(tCase.key, tCase.value)
			}
		})
	}
}

func BenchmarkGet(b *testing.B) {
	for _, tCase := range cases_for_get {
		b.Run(tCase.key, func(bb *testing.B) {
			s, _ := NewStorage()

			bb.ResetTimer()

			for i := 0; i < bb.N; i++ {
				s.Get(tCase.key)
			}
		})
	}
}

func BenchmarkSetGet(b *testing.B) {
	for _, tCase := range cases_for_setget {
		b.Run(tCase.key, func(bb *testing.B) {
			s, _ := NewStorage()

			bb.ResetTimer()

			for i := 0; i < bb.N; i++ {
				s.Get(tCase.key)
				s.Set(tCase.key, tCase.value)
			}
		})
	}
}

type setStructure struct {
	key   string
	value string
	count int
}

var cases_for_set = []setStructure{
	{"keySet1", "1", 5},
	{"keySet2", "2", 5},
	{"keySet3", "3", 5},
}

var cases_for_get = []setStructure{
	{"keyGet1", "4", 5},
	{"keyGet2", "5", 5},
	{"keyGet3", "6", 5},
}

var cases_for_setget = []setStructure{
	{"keySetGet1", "7", 5},
	{"keySetGet2", "8", 5},
	{"keySetGet3", "9", 5},
}
