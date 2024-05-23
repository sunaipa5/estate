package main

import (
	"estate/options"
	"fmt"
)

func main(){
	fmt.Println("Estate server running on:",options.Port)
	router()
}