package main

import (
	"HashTable/Model/Hash"
	"fmt"
)

func main() {
	h := Hash.New()
	(*h).Insert("one")
	(*h).Insert("two")
	(*h).Insert("three")
	(*h).Insert("four")
	(*h).Insert("test4")
	(*h).Insert("mohsen")
	(*h).Insert("Hello")
	fmt.Println((*h).Search("one"))
	fmt.Println((*h).Search("test"))
	fmt.Println((*h).Delete("mohsen"))

}
