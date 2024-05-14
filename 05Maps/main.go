package main

import "fmt"
func main() {
	fmt.Println("maps practise");
	var ma map[int]string;
	fmt.Println(ma); //attempts to write to a nil map will cause a panic runtime
	//initialise with make 
	m := make(map[int]string);
	m[1]="poojitha";
	m[2]="swarupa";
	m[3]="raju";
	m[4]="Lpgith vardhan";
	fmt.Println("map is",m)
	fmt.Println(m[1]);
	// If the requested key doesn’t exist, we get the value type’s zero value.
	fmt.Println(m[9]);
	fmt.Println(len(m));
	// The delete function doesn’t return anything, and will do nothing if the specified key doesn’t exist.
	delete(m,1);
	// A two-value assignment tests for the existence of a key:
	i,ok := m[1]
	fmt.Println("after searching",i,ok);
	// iterate
	fmt.Println("iteratio of map")
	for key,value := range m {
		fmt.Println("Key:", key, "Value:", value)
	}

	
	
}