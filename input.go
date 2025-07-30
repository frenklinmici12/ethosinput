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
        var errMsg kernelTypes.String
        errMsg = "Failed to read input"
        altEthos.Write("/programs/output", &errMsg)
        return
    }

    altEthos.Write("/programs/output", &input)
}
