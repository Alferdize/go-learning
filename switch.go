package main

import "fmt"

func main(){
var grade string = "B"
var marks float32 = 95.0

switch marks {
case 90: grade = "A"
case 80: grade = "B" 
case 50,60,70 : grade = "C"
default : grade = "D"
}
switch{
case grade == "A" :
fmt.Println("Excellent!\n")
case grade == "B" :
fmt.Println("Well Done!\n")
case grade == "C" :
fmt.Println("Good!\n")
case grade == "D" :
fmt.Println("You passed!\n")
case grade == "F" :
fmt.Println("Better try again!\n")
default :
fmt.Println("Invalid Grade!\n")
}
fmt.Println("Your Grade is",grade)
}