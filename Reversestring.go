package main

import "fmt"

func reversestr(str string) string {
	runes := []rune(str)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)

}

func main() {
	var input string

	fmt.Println("Enter the string : ")
	fmt.Scan(&input)

	reversed := reversestr(input)
	fmt.Println("The reversed string: ", reversed)
}
