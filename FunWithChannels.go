package main

import (
	"fmt"
	"time"
)


func readWriteToChannel(tube chan string) {
	fmt.Println("Inside readWriteToChannel method - START")
	time.Sleep(10 * time.Second)
	tube <- "It's a brand new day"
	time.Sleep(10 * time.Second)
	//fmt.Println("Inside readWriteToChannel method - END")

}

func main(){
	fmt.Println(" Fun begins now...")
	myPipe := make(chan string)
	go readWriteToChannel(myPipe)
	fmt.Println("Got response from channel ", <- myPipe)
}