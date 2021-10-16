package main

import (
	"fmt";
	"encoding/json";
)

type vehicle interface {
	start() string
}

type Car struct {
	CarName string `json:"car"`
	CarYear int `json:"year"`
}

func (c Car) drive() {
	fmt.Println(c.CarName, "acelerou!")
}

func (c Car) start() string {
	return "started"
}

func main() {
	var i int = 42
	fmt.Println(i)

	var n bool = true
	fmt.Println(n)

	ab := [...]int{1,2,3}
	ac := [...]int{1,2,3}
	b := ab
	c := &ac

	fmt.Println("Begin:")
	fmt.Printf("a: %v \nb: %v \nc: %v\n", ab,b,c)

	fmt.Println("Change b[1] to 5:")
	b[1] = 5
	fmt.Printf("a: %v \nb: %v \n", ab,b)

	fmt.Println("Change c[1] to 5:")
	c[1] = 5
	fmt.Printf("a: %v \nc: %v\n", ac,c)

	fmt.Println("Slices")

	list := []int{1,2,3,4,5,6,7,8,9,10}
	listB := list[:]
	listC := list[3:]
	listD := list[:6]
	listE := list[3:6]
	fmt.Println(list)
	fmt.Println(listB)
	fmt.Println(listC)
	fmt.Println(listD)
	fmt.Println(listE)


	// Some kind of strange classes
	var car1 Car = Car{
		CarName: "Golzinho",
		CarYear: 2013,
	}

	fmt.Println(car1)
	car1.drive()
	fmt.Println()

	fmt.Println("Test1")
	result, _ := json.Marshal(car1)
	fmt.Println(string(result))
	fmt.Println()
	
	fmt.Println("Test2")
	
	j := []byte(`{"car":"BMW","year":2021}`)
	
	var car2 Car
	json.Unmarshal(j, &car2)
	

	fmt.Println(car2.CarName)

}