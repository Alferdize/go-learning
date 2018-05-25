package main

/* The first line used to tell the compiler in which package the program should lie
as go program run in packages so it is mandatory to write packages .
The main package is the starting point of the program
Each package has a path and name associated with it*/


import "fmt"


/* The import "fmt" is a preprocessor command which tells the go compiler
to  include which files of package "fmt"*/
func main(){


	// This is the main functon where program execution begins


	fmt.Println("hello")


	//This is the another function of go which display the message "hello world"
	//on screen
	//Here "fmt" has imported the Println on the screen

	//In go if the name starts with capital letter it means it is exported
	//Exported means the function or variable/constant is
	//accessible to the importer of the respective package
}
