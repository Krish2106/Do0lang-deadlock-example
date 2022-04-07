package main

import (
	"fmt"
)

var paragraph = []string{"There are usually", "about 200 words"," in a paragraph", "but this can vary widely"}

func main(){

	comm := make(chan string)
	done := make(chan bool)

	go func() {
		prodWorker(comm)
	}()

	go consWorker(comm, done)

	<- done
}

func prodWorker( p chan string){

	for _,v := range paragraph{
		p <- v
	}
	close(p)
}

func consWorker(c chan string, done chan bool) {
	for v := range c{
		fmt.Println(v)
	}
       //Uncomment below line to release deadlock	
       //done <- true
}
