package main
import "fmt"


func main(){
	a:= 6
	changeValue(a)
	fmt.Printf("gia tri a after changeValue la: %d \n", a)

	changeValueByAddress(&a)

	fmt.Printf("gia tri a after changeValueByAddress la: %d \n", a)

}

func changeValue(a int){
	a = 100
}

func changeValueByAddress(a *int){
	*a = 200
}