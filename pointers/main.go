package main

import "fmt"

func main() {
	fmt.Println("welcome to poinetrs")
	var ptr *int
	fmt.Println("defaukt value of a pointer ",ptr);
	myNum := "poojitha"
	var ptr1 = &myNum
	fmt.Println("address of a pointer ",ptr1);
	fmt.Println("value of a pointer ",*ptr1);
	*ptr1 = *ptr1 + "reddy"
	fmt.Println(*ptr1);

}