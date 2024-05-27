package main

import (
	"fmt"
)

func main() {

	hobbies := [3]string{"gym", "swimming", "travel"}
	fmt.Println("My hobbies are: ", hobbies)

	fmt.Println("First Hobby:", hobbies[0])
	fmt.Println("Rest of the Hobbies:", hobbies[1:])

	task3 := hobbies[:2]
	fmt.Println("task3:", task3)

	task4 := task3[1:3]
	fmt.Println("task4:", task4)

	products := []*Product{NewProduct("p1", "id1", 34.87), NewProduct("p2", "id2", 98.9)}
	fmt.Println(products)
	products = append(products, NewProduct("p3", "id3", 87.0))
	fmt.Println(products)

	fmt.Println(add(5, 6))
	fmt.Println(add(5.8, 6.5))
	fmt.Println(add("Mo", "Ba"))
}

type Product struct {
	Title string
	Id    string
	Price float64
}

func NewProduct(title, id string, price float64) *Product {
	return &Product{
		title, id, price,
	}
}

func (p *Product) String() string {
	return fmt.Sprintln(p.Id, p.Title, p.Price)
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}
