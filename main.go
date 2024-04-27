package main

import(
"fmt"
)

func main(){
	fmt.Println("Estate server running on:",getListenAdress())
	router()
}