package main

import (
       "fmt"
       // "math/rand"
       //"math"
       "crypto/rand"
       "reflect"
       // "strconv"
       // "sort" // sorts numbers maybe?
       // "encoding/binary"
       //"bufio" // https://golang.org/pkg/bufio/
       //"os" // https://golang.org/pkg/os/
)

// PRE-DEFINED VARIABLES

var (
    // [var name]	[type]	= [value]
    DiSides	int // note: only assign type, not value
    DiNum	int = 10
)

func add(x , y int) int {
	return x + y
}

func multiply(x , y int) int {
	return x * y
}

func swap(x, y string) (string, string) {
	return y, x
}

func rangedown(randvalue, oldmax, newmax int) int {
	return ((randvalue - 1) * (newmax - 1)) / (oldmax - 1) + 1
}

func main() {
	// User Input Test (STRINGS)
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Print("Enter text: ")
	//text, _ := reader.ReadString('\n')
	//fmt.Println(text)
	
	// Import Interger for DiSides variable
	fmt.Println("How many Sides on the Di? ")
	fmt.Scanf("%d", &DiSides)
	
	// Generate Random string of bytes
	// kudos to https://flaviocopes.com/go-random/
	// for crypto rand generation
	key := [256]byte{}
    _, err := rand.Read(key[:])
    if err != nil {
        panic(err)
    }
    
    // Assign Random String to variable RandNumString
    // RandNumString := key
    
    // Other "How does golang work" shit.
	//fmt.Printf("%d dice with %v sides per di\n",DiNum, DiSides)
	//fmt.Printf("Highest possible roll = %v\n",multiply(DiNum, DiSides))
	//fmt.Printf("This is the result of math.Pi function: %v\n", math.Pi)
	//SidesPlusNum := add(DiNum,DiSides) // define var, w/custom function
	//fmt.Printf("Sides + Number of Dice = %v\n", SidesPlusNum)
	//a, b := "HELLO", "WORLD"
	//fmt.Println(swap(a, b))
	
	// print value of 1st number from random generator
	fmt.Printf("%v\tThe original number\n", key[0])
	// view variable type
	fmt.Printf("%v\tThe original var type\n", reflect.TypeOf(key[0])) 
    // convert key to string
    TestRoll := int(key[0]) 
    fmt.Printf("%v\tThe converted interger roll\n",TestRoll)
    // custom convert function to number of Di Sides
    ThisRoll := rangedown(TestRoll,256,DiSides)
	fmt.Printf("%v\tUnrounded conv. to %v sides\n",ThisRoll,DiSides)

}
