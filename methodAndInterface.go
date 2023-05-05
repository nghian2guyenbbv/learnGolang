package main
 import ("fmt")
 type Square struct{
	X float64
 }

 func (s Square) Acreage() float64{
	return s.X * s.X
 }

 type Retangle struct{
	X float64
	Y float64
 }
 func (r Retangle) Acreage() float64{
	return r.X * r.Y
 }


 func main(){
	s:=Square{4}
	fmt.Println("dien tich hinh vuong la ", s.Acreage())
	r:=Retangle{5, 6}
	fmt.Println("dien tich hinh chu nhat la: ", r.Acreage());
 }