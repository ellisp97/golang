package main

import (
	"fmt"
	"strings"
)

func main(){
	var input string
	fmt.Scanf("%s\n", &input)
	answer := 1

	for _,char := range input {
		currLetter := string(char)
		if strings.ToUpper(currLetter) == currLetter {
			answer++
		}
	}
	fmt.Println(answer)
}