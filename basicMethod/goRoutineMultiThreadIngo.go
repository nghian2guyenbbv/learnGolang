package main
import ("fmt"
"time")

func main(){
	go print1to10()
	go print11to20()
	time.Sleep(time.Second)
	 /* print1to10()
	 print11to20() */
}

func print1to10(){
	for i := 0; i<=10; i++ {
		go fmt.Println(i)
	}
}

func print11to20(){
	for i := 11; i<=20; i++ {
		go fmt.Println(i)
	}
}