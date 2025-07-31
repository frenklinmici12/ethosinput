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
    status := altEthos.ReadStream(syscall.StdIn, &input)
    if status != syscall.StatusOk {
        var errMsg kernelTypes.String = "Failed to read from stdin"
        altEthos.WriteStream(syscall.StdOut, &errMsg)
        return
    }

    // Write back to stdout
    altEthos.WriteStream(syscall.StdOut, &input)

	log.Println("log: all done!!!!")
}
