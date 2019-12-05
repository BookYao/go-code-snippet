package main

import (
	"fmt"
	"unsafe"
)

func checkEncianness() {
	var numInt int32 = 0x01020304
	pInt := unsafe.Pointer(&numInt)
	pBin := (*byte)(pInt)

	if *pBin == 0x4 {
		fmt.Println("Little Endian")
	} else {
		fmt.Println("Big Endian")
	}
}

func main() {
	checkEncianness()
}
