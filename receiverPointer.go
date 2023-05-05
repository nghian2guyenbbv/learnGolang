package main

import ("fmt")
type Cordinates struct{
	X, Y float64
}
func (p *Cordinates)updateCoordinate(dx, dy float64){
	p.X = p.X + dx
	p.Y = p.Y + dy
}

func main(){
	p:= Cordinates{3, 4}
	p.updateCoordinate(7,9)
	fmt.Println("gia tri toa do moi la: ", p)
}