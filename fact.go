package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"time"
)

func Usage(name string) {
	fmt.Fprintf(os.Stderr, "usage: %s number\n", name)
}

func main() {
	if len(os.Args) != 2 {
		Usage(os.Args[0])
		os.Exit(1)
	}

	n, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	var f big.Int
	t := time.Now()
	fmt.Printf("%d! = %v\n", n, f.MulRange(1, n))
	fmt.Printf("time = %v\n", time.Since(t))
}
