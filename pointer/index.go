package pointer

import "fmt"

func PointerToPointer() {

	var x = 100

	fmt.Println("Value of x :", x)
	var ptr1 = &x
	fmt.Println("Value of ptr1", ptr1)
	var ptr2 = &ptr1
	fmt.Println("value of Address at ptr2: *ptr2", *ptr2)
	fmt.Println("*(value of address of ptr2)", **ptr2)

}
