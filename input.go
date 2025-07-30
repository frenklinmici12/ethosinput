package main

import (
    "ethos/syscall"
    "ethos/kernelTypes"
    "ethos/altEthos"
)

func main() {
    var input kernelTypes.String

    // Read user input from the console
    status := altEthos.Read("/programs/input", &input)
    if status != syscall.StatusOk {
        var errMsg kernelTypes.String
        errMsg.Value = "Failed to read input"
        altEthos.Write("/programs/output", &errMsg)
        return
    }

    // Echo input back to the output
    altEthos.Write("/programs/output", &input)
}
