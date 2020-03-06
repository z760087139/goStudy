package main

import "fmt"

type Person struct {
	name    string
	age     int
	friends []Person
	// friend  Person
}

func f1() {
	a := new(Person)
	fmt.Printf("%#v\n", a)
	b := Person{
		name: "B",
		age:  10,
		// friends: nil,
		// friends: []Person{
		// 	Person{
		// 		name: "b",
		// 		age:  1,
		// 	},
		// },
	}
	fmt.Printf("%#v\n", b)
	fmt.Println(b.name)
	b.name = "b"
	b.friends = []Person{
		Person{
			name: "b",
			age:  1,
		},
		Person{
			"S",
			2,
			nil,
		},
	}
	fmt.Printf("%#v\n", b)

}

func f2() {
	ps := Person{
		name: "a",
		age:  10,
	}
}

func f3(p Person) {
	p.name = "b"
	fmt.Println(p)
	fmt.Printf("%p\n", p)
}

func f4() {}

func main() {
	f1()
}
