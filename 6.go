package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	persons := make([]Person, 0)
	for i := 0; i < 5; i++ {
		tperson := Person{
			Name: fmt.Sprintf("gong:%d", i),
			Age:  i,
		}
		persons = append(persons, tperson)
	}
	for k, v := range persons {
		fmt.Println(k, v)
	}

}
