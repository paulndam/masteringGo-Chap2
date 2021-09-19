package main

import (
	"fmt"
	"time"
)


func runLoopSend(n int, ch chan int){
	for i := 0; i < n; i++{
		ch <- i
	}
	close(ch)
}


func runLoopRecieve(ch chan int){
	for i := range ch {
		fmt.Println("Recieved the value ---->",i)
	}
}


func main(){
	myChanel := make(chan int)
	go runLoopSend(20, myChanel)
	go runLoopRecieve(myChanel)

	time.Sleep(2 * time.Second)

	// creating buffered chanel.
	// myBufChan := make(chan int, 10)
}