package main

import(
	"fmt" //for input-output
)

func main(){
	var num int = 1

	var str string = "javier.jentzenpaolo@gmail.com"

	concatenated := fmt.Sprint(num, str)

	fmt.Println(concatenated)
}
