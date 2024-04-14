package main

import "fmt"
import "time"

func firstFunc() {
	fmt.Println("Kira7 a7ssen artist fi maruecos")
	fmt.Println("3aaaasseeeemaaaa")
}

func multiply(x int, y int) int {
	return x*y
}

func returnTwo(x int) (int, int) {
	return x*2,x/2 
}

func deferTesting() {
	defer fmt.Println("Fikom a khouya chi f3ayl d nissa2")
	defer fmt.Println("Ntiya ra qtita Siamois")
	fmt.Println("Kiraaaaaa")
	fmt.Println("Chno bghitit nclipi rap")
}

func main() {

	// First Go statements
	fmt.Println("First print statement by Go.")
	fmt.Println("fmt stands for format.")
	fmt.Println("Don't code before reading the docs.")
	fmt.Print("Between", "these", "two", "methods?")
  	fmt.Print("Anything odd about", "the spacing? \n")
	fmt.Println("Today time is :",time.Now())

	const funFact = "Go and awesome."
	const pi = 3.14 
	const runTime = 459834

	var isConnected bool
	var goals int8 = 5
	var price float32 = 456.45
	isConnected = true
	var day int = 59
	fmt.Println(isConnected,goals,price,day)
	
	var phrase string = "This is my phrase"
	fmt.Println(phrase+ " gol ghi hola kon ghi poli maradonna haz napoli")
	fmt.Printf("Here it is %v",phrase)

	// Zero Values
	var emptyInt int8
  	var emptyFloat float32
  	var emptyString string 
	var emptyBool bool
  	fmt.Println(emptyInt, emptyBool, emptyFloat, emptyString)

	// Inferring type
	lolPhrase := "is this serious"
	fmt.Println(lolPhrase)

	var numberOfHits = 435 
	fmt.Println(numberOfHits)

	// Multiple declarations
	number, quote := 456, "yes Go"
	fmt.Println(number,quote)

	tki, lop := "fiha", "mafiha"
	if (tki == lop) {
		fmt.Println("BAYNA FIHA")
	} else {
		fmt.Println("BAYNA MAFIHACH")
	}

	firstFunc()

	fmt.Println("The multiplication goes :",multiply(56,7))

	fmt.Println(returnTwo(6))

	deferTesting()

	// Adresses 
	x := "My very first address"
	fmt.Println(&x) 

}