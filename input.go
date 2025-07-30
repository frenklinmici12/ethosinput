package main

import (
    "ethos/altEthos"
    "ethos/fmt"
    "ethos/syscall"
)

func main() {
    // Open the input stream (user input from console)
    fd, status := altEthos.OpenReadStream("/programs/input")
    if status != syscall.StatusOk {
        fmt.Printf("Failed to open input: %v\n", status)
        return
    }

    var input string

    // Read from the input stream
    status = altEthos.ReadStream(fd, &input)
    if status != syscall.StatusOk {
        fmt.Printf("Failed to read: %v\n", status)
        return
    }

    // Echo the input back to the user
    fmt.Printf("You said: %v\n", input)
}
