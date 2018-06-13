package main

import (
       "fmt"
       "math/rand"
       "time"
       "io/ioutil"
       //"strings"
)

// PRE-DEFINED VARIABLES
var (
    // [var name]	[type]	= [value]
    DiSides	int // note: only assign type, not value
    DiNum	int // 
)

// Generate Random Number between Two Values
// via: https://golangcode.com/generate-random-numbers/
func random(min int, max int) int {
	max = max + 1 // ensures you can hit intended max number
    return rand.Intn(max-min) + min
}

// read file: ./eff_large_wordlist.txt

func main() {
	// Import Interger for DiSides variable
	fmt.Println("How many Sides on the Di? ")
	fmt.Scanf("%d", &DiSides)
	fmt.Printf("How many dice %v-sided to roll? ",DiSides)
	fmt.Scanf("%d", &DiNum)
	
	// seed random
	RollNum := 1
	for RollNum <= DiNum {
		rand.Seed(time.Now().UnixNano())
		randomNum := random(1, DiSides)
		fmt.Printf("%v:\t%d\n", RollNum, randomNum)
		RollNum = RollNum + 1
	}
	
	// Read File
	b, err := ioutil.ReadFile("./eff_large_wordlist.txt")
    if err != nil {
        fmt.Print(err)
    }
    EFF_Wordlist := string(b)
    fmt.Println(EFF_Wordlist)
    
    //for i in EFF_Wordlist {
		//if stings.Contains("11111", i) == TRUE {
			//fmt.Println(i)
		//}
}
