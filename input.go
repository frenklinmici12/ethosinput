package main

import (
    "ethos/altEthos"
    "ethos/fmt"
    "ethos/syscall"
)

func main() {
    var inputString string

    // Read user input from standard input (terminal)
    status := altEthos.ReadStream(altEthos.StdinFd, &inputString)
    if status != syscall.StatusOk {
        fmt.Printf("Failed to read input: %v\n", status)
        return
    }

    // Echo it back to the user
    fmt.Printf("You entered: %v\n", inputString)
}
