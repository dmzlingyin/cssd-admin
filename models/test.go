package models

import (
	"log"
)

type Test struct {
	ID   int
	Name string
}

func TestCreate() bool {
	test := &Test{ID: 3, Name: "Lv Wenchao"}
	result := db.Create(test)

	if result.Error != nil {
		log.Println("Fail to create record")
		return false
	}

	return true
}

func TestSelect() []Test {
	var test []Test
	db.Find(&test)

	return test
}
