package main

import (
	"fmt"
	//"strings"
)

func main(){
	var length, shift int
	var input string
	fmt.Scanf("%d\n", &length)
	fmt.Scanf("%s\n", &input)
	fmt.Scanf("%d\n", &shift)

	// ============= Method 1 =================

	//alphaHiString := "abcdefghijklmnopqrstuvwxyz"
	//alphaLowString := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	//ret := ""
	//for _, char := range input {
	//	switch {
	//	case strings.IndexRune(alphaLowString, char) >= 0:
	//		ret += string(shiftLetters(char, shift, []rune(alphaLowString)))
	//	case strings.IndexRune(alphaHiString, char) >= 0:
	//		ret += string(shiftLetters(char, shift, []rune(alphaHiString)))
	//	default:
	//		ret = ret + string(char)
	//	}
	//}

	// ============= Method 2 ==============

	var ret []rune
	for _,char := range input {
		ret = append(ret, cipher(char,shift))
	}

	fmt.Println(string(ret))
}

func cipher(r rune, shift int) rune {
	if r >= 'A' && r <= 'Z' {
		return shiftLetters(r, 'A',shift)
	}
	if r >= 'a' && r <= 'z' {
		return shiftLetters(r, 'a',shift)
	}
	return r
}

func shiftLetters(r rune, base, shiftSize int) rune{
	temp := int(r) - base
	temp = (temp + shiftSize) % 26 //size of alphabet
	return rune(temp + base)
}

//func shiftLetters(s rune, shiftSize int, key []rune) rune {
//	//index := -1
//	//for i, r := range key{
//	//	if r == s{
//	//		index = i
//	//		break
//	//	}
//	//}
//
//	//alternative go version to above loop
//	index := strings.IndexRune(string(key), s)
//
//	if index < 0 {
//		panic("havent found letter in alphabet")
//	}
//	index = (index + shiftSize) % len(key)
//	return key[index]
//}

