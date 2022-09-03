package main

import(
	"fmt" //for input-output
	"math" //for math class
	"errors" //for error class
)

func main(){
	result := sum(2, 3)
	fmt.Println(result)

	sq, err := sqrt(16)
	fmt.Println(sq)
	//catches error for negative numbers
	if err != nil{
		fmt.Println(err)
	} else{
		fmt.Println(sq)
	}
}

func sum(num01 int, num02 int) int{
	return num01 + num02
}

func sqrt(num01 float64) (float64, error){
	if num01 < 0 { //catches error for negative numbers
		return 0, errors.New("Undefined for negative numbers")
	}
	return math.Sqrt(num01), nil
}

