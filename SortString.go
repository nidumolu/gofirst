package main

import  (
	"fmt"
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

func main(){
   str1 := "sortthisstring"  
   s := []rune(str1)
   sort.Sort(sortRunes(s))
   fmt.Println(string(s))
}