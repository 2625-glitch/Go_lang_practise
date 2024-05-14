package main

import "fmt"
func  main()  {
	fmt.Println("maps practise");
	user1 := User{"poojitha",21,true};
	fmt.Println(user1);
	fmt.Printf("user1 details are: %+v\n",user1);
	fmt.Printf("specific details are name: %v and age %v",user1.Name,user1.Age);
}
type User struct {
	Name string
	Age int 
	Status bool
}