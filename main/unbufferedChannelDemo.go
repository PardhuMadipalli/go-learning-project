package main

import (
	"fmt"
	"time"
)

func main() {
	// Unbuffered channel
	fmt.Println("Started main func")
	ch := make(chan string)
	go sendRoutine(ch)
	timeAfterStartingMain := time.Now().UnixMilli()
	messageReceived := <-ch
	fmt.Printf("Message received from channel in main: %v after %v milli seconds.\n", messageReceived, time.Now().UnixMilli()-timeAfterStartingMain)
	timeAfterStartingMain = time.Now().UnixMilli()
	messageReceived = <-ch
	fmt.Printf("2nd Message received from channel in main: %v after %v milli seconds.\n", messageReceived, time.Now().UnixMilli()-timeAfterStartingMain)
}

func sendRoutine(ch chan<- string) {
	fmt.Println("Starting routine sendRoutine")
	fmt.Println("sendRoutine waiting for 5 seconds")
	time.Sleep(5 * time.Second)
	ch <- "Hello world!"
	fmt.Println("sendRoutine waiting for 3 seconds")
	time.Sleep(3 * time.Second)
	ch <- "hello world for 2nd time!"
	fmt.Println("Sent hello world to the channel.")
}
