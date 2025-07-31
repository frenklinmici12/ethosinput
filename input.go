package main

import (
	"ethos/altEthos"
	"ethos/fmt"
	"ethos/kernelTypes"
	"ethos/syscall"
	

	"log"
)

func main() {
	
    var input kernelTypes.String

	fmt.Printf("enter text: ")

    // Read from stdin (0) stdout is 1
    status := altEthos.ReadStream(0, &input)
    if status != syscall.StatusOk {
        var errMsg kernelTypes.String = "Failed to read from stdin"
        altEthos.WriteStream(1, &errMsg)
        return
    }

	fmt.Println(input)

	log.Println("log: all done!!!!")
	
}
