package main

import(
	"fmt" //for input-output
)

type person struct{
	name string
	age int
}

func main(){
	p := person{name: "Jentzen Paolo", age: 24}
	fmt.Println(p)
	fmt.Println(p.name)
	fmt.Println(p.age)
}
