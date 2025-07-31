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

    // Read from stdin
    status := altEthos.ReadStream(0, &input)
    if status != syscall.StatusOk {
        var errMsg kernelTypes.String = "Failed to read from stdin"
        altEthos.WriteStream(1, &errMsg)
        return
    }

    // Write back to stdout
    //altEthos.WriteStream(1, &input)
	fmt.Println(input)

	log.Println("log: all done!!!!")
	
}
