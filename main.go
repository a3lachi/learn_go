package main

import "fmt"
import "time"


func firstFunc() {
	fmt.Println("Kira7 a7ssen artist fi maruecos")
	fmt.Println("3aaaasseeeemaaa")
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

type Cordinate struct {
	x int 
	y int
}

func friCords(cords Cordinate) int {
	return cords.x * cords.y
}

func (cords Cordinate) area() int {
	return cords.x * cords.y
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

	// Pointers
	var strPointer *string 
	strPointer = &x
	fmt.Println(strPointer) 
	fmt.Println(*strPointer) 

	// For loops
	for counter:=0;counter<10;counter++ {
		fmt.Print(counter," ")
	}
	fmt.Println("")

	for count := 1; count <= 12; count=count+2 {
		fmt.Print(count," ")
	}
	fmt.Println("")

	// While Loops
	whileCounter:=1
	for whileCounter<20 {
		fmt.Print(whileCounter,"-")
		whileCounter++
	}
	fmt.Println("")


	// Arrays & Slices
	var intArray []int // [5]int
	intArray = []int{5,6,7,8,9,10}
	// intArray = intArray[1:3]
	intArray = intArray[:len(intArray)-1]

	floatArray := [5]float32{1.1,2.2,3.3,4.4}
	zeroes := [3]float64{}
	fmt.Println(intArray,floatArray,zeroes)

	for index, value := range floatArray {
		fmt.Println(index,value)
	}


	// Maps
	newMap := map[string]int{
		"simo":667,
		"mohamed":133,
		"kira7":555,
	}
	fmt.Println("New Map :",newMap)
	fmt.Println("First Map Eelement :",newMap["simo"])

	newMap["toto"]=567
	newMap["khoo"]=212

	delete(newMap,"simo")

	for key, value := range newMap {
		fmt.Printf("The key %v and the value %v\n",key, value)
	}
	
	names := []string{"Kathryn", "Martin", "Sasha", "Steven"}
	fmt.Println("Names length :",len(names))

	books := []string{"Tom Sawyer", "Of Mice and Men"}
	books = append(books, "Frankenstein")
	books = append(books, "Dracula")
	fmt.Println(books)


	// Maps
	mymap := map[string]int{}
	fmt.Println(mymap)


	// Structs 
	cords := Cordinate{69,23}
	fmt.Println(cords.area(),cords,friCords(cords))

}