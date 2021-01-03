package main

//#include <stdio.h>
//#include <strings.h>
//#include <stdlib.h>
/*

unsigned char array[2];

unsigned char *foo() {
    array[0] = 0x1;
    array[1] = 0x2;
    return array;
}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	p := C.foo()
	fmt.Println(p)

	up := unsafe.Pointer(p)
	fmt.Println(up)

	b := C.GoBytes(up, 2)
	fmt.Println(b)
}
