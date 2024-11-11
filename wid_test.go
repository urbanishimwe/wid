package wid

import (
	"fmt"
	"testing"
)

const generateCount = 1000

func TestGenerate(t *testing.T) {
	uniqueMap := map[string]bool{}
	for i := 0; i < generateCount; i++ {
		id, err := Generate()
		if err != nil {
			t.Fatal(err)
		}
		if uniqueMap[id] {
			t.Fatal("unique ID collision")
		}
		uniqueMap[id] = true
		fmt.Println(id)
	}
}

func TestGenerateUpper(t *testing.T) {
	uniqueMap := map[string]bool{}
	for i := 0; i < generateCount; i++ {
		id, err := GenerateUpper()
		if err != nil {
			t.Fatal(err)
		}
		if uniqueMap[id] {
			t.Fatal("unique ID collision")
		}
		uniqueMap[id] = true
		fmt.Println(id)
	}
}
