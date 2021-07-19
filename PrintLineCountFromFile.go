//  prints the text of each line that appears more than
// once in the standard input read fron a file passed as argument, preceded by its count.
// go run PrintLineCountFromFile.go test.txt
package main

import (
    "bufio"
    "fmt"
    "os"
	"time"
	"log"
)

func main() {
    counts := make(map[string]int)
    /*input,err := bufio.NewScanner(os.Args[1])
    for input.Scan() {
        counts[input.Text()]++
    }
    */
    if os.Args != nil && len(os.Args) > 1 {
	f , err := os.Open(os.Args[1])
		if err != nil  {
			err := fmt.Errorf("error occurred at: %v", time.Now())
			log.Fatal(err)
		}
	
	
	input := bufio.NewScanner(f)
    for input.Scan() {
        counts[input.Text()]++
    }

    for line, n := range counts {
        if n > 0 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
    
} else {
	err1 := fmt.Errorf("No file to parse path has been provided: Curr time", time.Now())
	log.Fatal(err1)

}
}