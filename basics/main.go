package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)
const Token string = "ex of const variables by default public"
func main() {
	fmt.Println("Hey..i just started learningb go with my love after making out and release of energy");
	var username string = "poojithareddy";
	fmt.Println(username)
	fmt.Printf("Variable is of type %T ]n ", username);
	var exbool bool = true;
	fmt.Println(exbool);
	fmt.Printf("Variable is of type %T \n", exbool);
	//1.  implicit and no var type variables
	var eximplicit = "Example of implicit type";
	fmt.Println(eximplicit);
	exnovar := "example of no var";
	fmt.Println(exnovar); // allowed only insided methods 
	fmt.Println(Token);


	//2. user Input 
	reader := bufio.NewReader(os.Stdin);
	fmt.Println("enter rating for my work");
	// comma ok syntax or error ok
	input, _ := reader.ReadString('\n');
	fmt.Println("Output is," ,input);
	fmt.Printf("type of taken user input is %T\n",input);

	//3.conversion of user input strings and numbers
	converted, err := strconv.ParseFloat(strings.TrimSpace(input),64);
	if err!=nil {
		fmt.Println(err);
	}else{
		fmt.Println("after converting input ,",converted+1);
	}

	//4.time in golang
	presentTime := time.Now();
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("02-01-2006 15:04:05 Monday"));
	createdDate := time.Date(2024,time.October,05,18,00,00,00,time.UTC);
	fmt.Println(createdDate);
}