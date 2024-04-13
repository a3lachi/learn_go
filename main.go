package main

import "fmt"
import "time"

func main() {

	// First Go statements
	fmt.Println("First print statement by Go.")
	fmt.Println("fmt stands for format.")
	fmt.Println("Don't code before reading the docs.")
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

	var emptyInt int8
  	var emptyFloat float32
  	var emptyString string 
	var emptyBool bool
  	fmt.Println(emptyInt, emptyBool, emptyFloat, emptyString)

	lolPhrase := "is this serious"
	fmt.Println(lolPhrase)

	var numberOfHits = 435 
	fmt.Println(numberOfHits)


}