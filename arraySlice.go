package main

import (

"fmt"

)

func main(){

	number := [6]int{2, 4, 5, 7,8, 10 }

	// khoi tao slice

	var s []int = number[2:5]

	// In ra gia tri slice

	fmt.Println(s)

}