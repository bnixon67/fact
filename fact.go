package main

import (
"fmt"
"math/big"
"time"
)

func main() {
const n = 10000
var f big.Int
t := time.Now()
fmt.Printf("%d! = %v\n", 5, f.MulRange(1, 5))
fmt.Printf("%d! = %v\n", 100, f.MulRange(1, 100))
fmt.Printf("%d! = %v\n", n, f.MulRange(1, n))
fmt.Printf("time = %vs\n", time.Now().Sub(t).Seconds())
}
