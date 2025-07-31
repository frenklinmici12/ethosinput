package main

import (
	"ethos/altEthos"
	"ethos/fmt"
	"ethos/kernelTypes"
	"ethos/syscall"
	"fmt"

	"log"
)

func main() {
	
    var input kernelTypes.String

    // Read from stdin
   /* status := altEthos.ReadStream(0, &input)
    if status != syscall.StatusOk {
        var errMsg kernelTypes.String = "Failed to read from stdin"
        altEthos.WriteStream(1, &errMsg)
        return
    }*/
	fmt.Printf("enter text: ")

	fmt.Scanf("%s", &input)


	fmt.Println(input)

	log.Println("log: all done!!!!")
	
}
