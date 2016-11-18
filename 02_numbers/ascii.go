package main

import "fmt" 

func main() {
	for i := 1000 ;i < 2000 ; i++ {
		fmt.Printf("%d \t %b \t %X \t %q \n",i,i,i,i)
	}
}
