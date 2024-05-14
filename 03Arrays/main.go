package main

import "fmt"
func main() {
	var arr [4]string;
	arr[0]="swarupa"
	arr[1]="raju"
	arr[2]="lohithvardhan"
	fmt.Println("array is:",arr);
	fmt.Println("length is:",len(arr));
	var arr1 = [3]int{1,2}
	fmt.Println("array2 is:",arr1);
	for i:=0; i<len(arr1);i++ {
		arr1[i] = arr1[i] * 2;
	}
	fmt.Println(arr1);
}