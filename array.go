package main

import(
	"fmt" //for input-output
)

func main(){
	var numArray01 [5]int //first way to declare an array

	numArray02 := [5]int{0, 1, 2, 3, 4} //second way to declare an array

	numArray03 := []int{5, 6, 7, 8, 9} //"slices array" third way to declare an array without size
	numArray03 = append(numArray03, 10)

	strArray := []string{"a", "b", "c"}

	fmt.Println(numArray01)
	fmt.Println(numArray02)
	fmt.Println(numArray03)
	fmt.Println(strArray)
}
