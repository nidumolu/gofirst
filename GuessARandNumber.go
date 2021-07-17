package main

import
( 
	"fmt"
	"bufio"
	"os"
	"log"
	"time"
	"math/rand"
	"strconv"
	"strings"
)

func main(){
	seed := time.Now().Unix()
	rand.Seed(seed)
	numToGuess := rand.Intn(100) + 1

	//fmt.Println(" Random num =", numToGuess)
	reader := bufio.NewReader(os.Stdin)
	for guessCounter :=0 ; guessCounter < 10 ; guessCounter++ {
	fmt.Println("Guess a number bewtween 1 to 100..you have ..",10-guessCounter," left" )
	input , err := reader.ReadString('\n')
	if(err != nil){log.Fatal(err)}

	input = strings.TrimSpace(input)
	inputInt,_ := strconv.Atoi(input)
	fmt.Println("You entered..",input)

	if inputInt == numToGuess { 
		fmt.Println("You guessed it correct !") 
		break
	} else if inputInt < numToGuess { fmt.Println("You guess was lower !")
	} else if inputInt > numToGuess { fmt.Println("You guessed was higher!")
	} 
	
	if (guessCounter == 9) {
		fmt.Println("Sorry ..you ran out of chances..")
	}

	}

}