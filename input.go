package main

import (
    "ethos/syscall"
    "ethos/kernelTypes"
    "ethos/altEthos"
)

func main() {
    var input kernelTypes.String
    status := altEthos.Read("/programs/input", &input)
    if status != syscall.StatusOk {
        altEthos.Write("/programs/output", &kernelTypes.String{"Failed to read input"})
        return
    }

    // Echo the input back to output
    altEthos.Write("/programs/output", &input)
}
