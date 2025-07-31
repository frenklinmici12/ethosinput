package main

import (
    "ethos/syscall"
    "ethos/kernelTypes"
    "ethos/altEthos"
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
    altEthos.WriteStream(1, &input)

	log.Println("log: all done!!!!")
}
