package main
 import ("fmt")
 type Square struct{
	X float64
 }

 func (s Square) Acreage() float64{
	return s.X * s.X
 }

 func main(){
	s:=Square{4}
	fmt.Println("dien tich hinh vuong la ", s.Acreage())
 }