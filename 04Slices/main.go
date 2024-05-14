package main

import "fmt"
func main() {
	fmt.Println("slices practise");
	var arr = []int{1,2,3};
	fmt.Printf("type of arr: %T\n",arr);
	arr = append(arr,4,5);
	fmt.Println(arr);
	// remove 2 and store remaining elements in the same array
	arr = append(arr[:1],arr[2:]... );
	fmt.Println("after remmving 2 from the array",arr);
	// using make 
	a := make([]int, 5)
	fmt.Println(a);

}