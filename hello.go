package main

import (
    "fmt"
    "math/cmplx"
)

var (
    ToBe    bool        = false
    MaxInt  uint64      = 1<<64 - 1 
    c, java int         = 2, 3
    z       complex128  = cmplx.Sqrt(5-2i)
)

func split(sum int) (x,y int) {
    x = sum *4 / 9
    y = sum - x
    return x, y
}

func main() {
    var i = "9n"
    k := "tres"
	fmt.Println(split(14))
	fmt.Println(c, java, i, k)
	fmt.Printf("\nTipo %T, valor = %v\n", ToBe, ToBe)
	fmt.Printf("Tipo %T, valor = %v\n", MaxInt, MaxInt)
	fmt.Printf("Tipo %T, valor = %v\n", z, z)
}
