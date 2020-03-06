package main

import (
	"fmt"
)

func f1(){
	var m1 map[int]string
	m1 = map[int]string{}
	m1[0] = "A"
	m2 := make(map[string]int)
	m3 := map[int]string{}
	fmt.Println(m1,m2,m3)
	fmt.Println(m2["A"])
}
func f2(){
	set := make(map[string]bool)
	set["created"] =true
	fmt.Println(set["created"],set["none"])
	if _,ok := set["none"];!ok{
		fmt.Println("none not exists")
	}
}

func f3(){
	m := make(map[string]int,0)
	fmt.Printf("first :%p\n",m)
	m["A"] = 1
	fmt.Printf("second :%p\n",m)
	f4(m)
	fmt.Printf("fifth :%p\n",m)
	fmt.Println(m)
}

func f4(m map[string]int){
	fmt.Printf("third %p \n",m)
	m["A"] = 10
	fmt.Println(m["A"])
	m["B"] = 20
	fmt.Printf("fourth %p\n",m)
}

type Person struct{
	name string
}

func f5(){
	m := make(map[string]Person)
	m["M"] = Person{
		name :"M",
	}
	fmt.Println("f5 before f6",m)
	f6(m)
	fmt.Println("f5 after f6",m)
	m2 := make(map[string][]int)
	m2["M"] = []int{1,2,3}
	// fmt.Println(len(m2["M"]),cap(m2["M"]))
	f7(m2)
	fmt.Println(m2)
}

func f6(m map[string]Person){
	p := m["M"]
	p.name = "m"
	fmt.Println("f6 change1",m)
	m["M"] = p
	fmt.Println("f6 chagne2",m)
}

func f7(m map[string][]int){
	d := m["M"]
	// 修改d[0]的值是否对m产生影响
	d[0] = 0 
	fmt.Println(m)
	// 增加d的值是否对m产生影响
	d = append(d,4)
	fmt.Println(m)
}

func f8(){
	m := make(map[string][]int)
	m["m"] = []int{1,2,3}
	m["m"][0] = 0
	fmt.Println(m)
}

func f9(){
	m := make(map[string]Person)
	m["m"] = Person{
		name : "m",
	}
	// m["m"].name = "M"
	fmt.Println(m["m"].name)
}

func main(){
	f5()
}
