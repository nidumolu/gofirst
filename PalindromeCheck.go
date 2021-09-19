package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"log"
	"sort"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
    return s[i] < s[j]
}

func (s sortRunes) Len() int{
   return len(s)
}
func (s sortRunes) Swap(i, j int) {
   s[i], s[j] = s[j], s[i]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter two strings (with space inbetween) to check if they are palindromes, e.g cat act")	
	input,_ := reader.ReadString('\n')
	stringsArr := strings.Split(input, " ")
 
	if len(stringsArr) < 2 {
		log.Fatal(" You needed to enter two strings, bye.")
	}
	fmt.Println("You entered ", len(stringsArr), " words. ",stringsArr[0]," ", stringsArr[1])

	bool1 := checkPalindrome(strings.TrimSpace(stringsArr[0]), strings.TrimSpace(stringsArr[1]))
	fmt.Println(" And the result is :",bool1)
}

func checkPalindrome(arg1 string, arg2 string) bool {
	if(len(arg1) != len(arg2) ){
		log.Fatal(" Not a palindrome , lengths of two words are not equal ",arg1," ", arg2)
		return false
	}


   return checkIfRuneMapsareSame(arg1 , arg2 )
}

func checkIfRuneMapsareSame(arg1 string , arg2 string) bool {
	  
	s1 := []rune(arg1)
	sort.Sort(sortRunes(s1))

	s2 := []rune(arg2)
	sort.Sort(sortRunes(s2))

	if(string(s1) == string(s2)) {
		return true
	}
	return false
	
}
