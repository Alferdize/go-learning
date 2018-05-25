package main

import "fmt"

func getSequence() func() int {
	i := 0
	return func() int{
		i *= i
		return i
	}
}


func main(){
	/* nextMember is now a function with i as 0 */
	nextNumber := getSequence()
	/* invoke nextNumber to increase i by 1 return the same */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
}
