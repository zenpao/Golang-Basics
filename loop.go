package main

import(
	"fmt" //for input-output
)

func main(){
	//while loop
	x := 0
	for x < 5{
		fmt.Println(x)
		x++
	}

	//for loop
	for num:= 0; num < 5; num++{
		fmt.Println(num)
	}

	//array loop print
	strArray := []string{"a", "b", "c"} //string array
	for index, value := range strArray {
		fmt.Println("index", index, "value", value)
	}

	//map loop print
	maps := make(map[string]string) //string map

	maps["x"] = "one"
	maps["y"] = "two"
	maps["z"] = "three"

	for key, value := range maps {
		fmt.Println("key", key, "value", value)
	}
	
}
