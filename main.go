package main

import (
	"log"
	"os"
	"strings"

	"github.com/atotto/clipboard"
)

func main() {
    if len(os.Args) < 3 {
        return
    }
    val := strings.Join(os.Args[2:], " ")
    err := clipboard.WriteAll(val)
    if err != nil {
        log.Fatal(err)
    }
}
