package main

import "fmt"

func myfunc(ch chan string) {
	fmt.Println("saya" + <-ch)
}

func main() {
	fmt.Println("Start Main Method")
	ch := make(chan string)
	go myfunc(ch)
	// ch <- 23
	ch <- "alief"
	fmt.Println("End Main Method")
}
