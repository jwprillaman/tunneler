package main

import(
	"flag"
	"fmt"
)

func main(){

	var config = flag.String("c", "", "Configuation file to mimic.")
	flag.Parse()
	fmt.Printf("Using : %v\n", *config)

}
