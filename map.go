package main

import(
	"fmt" //for input-output
)

func main(){
	maps01 := make(map[string]int) //int map

	maps01["triangle"] = 0
	maps01["circle"] = 1
	maps01["square"] = 2

	//delete (maps01, "circle")

	maps02 := make(map[string]string) //string map

	maps02["x"] = "one"
	maps02["y"] = "two"
	maps02["z"] = "three"

	fmt.Println(maps01)
	fmt.Println(maps01["square"])
	fmt.Println(maps02)
	fmt.Println(maps02["z"])
}
