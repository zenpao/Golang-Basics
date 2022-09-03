package main

import(
	"fmt" //for input-output
	"strconv" //for string conversions
)

func main(){
	var num int = 1

	var str string = "javier.jentzenpaolo@gmail.com"

	converted2str := strconv.Itoa(num) //convert int to string

	converted2int, err := strconv.Atoi(str) //convert string to int


	fmt.Println(converted2str)

	 if err != nil {
        //handle error
	fmt.Println(converted2int)
	}
}
