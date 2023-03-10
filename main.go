package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func example(str []string) {
	fmt.Printf("%v, World!", str[0])
}

func main() {
	f := example
	ptrString := fmt.Sprintf("%d", &f)
	arr:=[]string{"HELLO"}
	runMe(ptrString, arr)
}
func runMe(ptrString string,args []string) {
	//Convert it to a uint64
	ptrInt, _ := strconv.ParseUint(ptrString, 10, 64)

	//They should match
	fmt.Printf("Address as String: %s as Int: %d\n", ptrString, ptrInt)

	//Convert the integer to a uintptr type
	ptrVal := uintptr(ptrInt)

	//Convert the uintptr to a Pointer type
	ptr := unsafe.Pointer(ptrVal)

	//Get the string pointer by address
	funcPtr := (*func([]string))(ptr)

	fmt.Printf("funcPtr %d Type=%T\n", funcPtr, funcPtr)

	//Run the function
	(*funcPtr)(args)
}
