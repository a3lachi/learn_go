

```
      `.-::::::-.`
  .:-::::::::::::::-:.
  `_:::    ::    :::_`
   .:( ^   :: ^   ):.
   `:::   (..)   :::.
   `:::::::UU:::::::`                   Exploring Go concepts.
   .::::::::::::::::.
   O::::::::::::::::O
   -::::::::::::::::-
   `::::::::::::::::`
    .::::::::::::::.
      oO:::::::Oo
```

# Go

Develeped by Google for in house software development, to make their software easier to write, maintain and build.

It became known in tech community and is used for writing APIs and DevOps tools.

Go sits in the middle between low-level languages and high-level languages, often offering the advantages of both.

## Compiling Files

To compile a file to an executable we use the command : 
```
git build myfile.go
```

To execute the file you call :
```
./myfile
```


## Running Files

To run directly a Go file use the command : 
```
go run myfile.go
```


## Go Doc

Go includes a program go doc for extracting and viewing documentation from .go files.
```
go doc fmt
go doc fmt.Println
```


## Values and Variables

How to store “data” by creating and using variables in Go.


### Const 

We use the const keyword to create a constant. We immediately assign a value to the constant using a literal.
```
const funcFact = "Go is awesome!"
fmt.Println(funFact)
```

### Numeric Types 

Go has 15 different numeric types that fall into the three categories: int, float, and complex.

```
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32
```


### Default int type

By providing the type int or uint, Go will check to see if the computer architecture is running on is 32-bit or 64-bit. Then it will either provide a 32-bit int (or uint) or a 64-bit one depending on the computer itself.

```
var timesWeWereFooled int
var foolishGamesPlayed uint
```



### Variable
Variables are defined with the var keyword and two pieces of information: the name that we will use to refer to them and the type of data stored in the variable. Since variables can be updated we don’t even need to assign a value initially.

```
var lengthOfSong uint16
var isMusicOver bool
var songRating float32
```

### String

A string is Go’s type for storing and processing text.

```
var phrase string
var logo string = "follow your dreams"
var description string
description = "This is our logo "+logo 
```


### Zero Values

Even before we assign anything to our variables they hold a value. Go’s designers attempted to make these “sensible defaults” that we can anticipate based on the variable’s types. All numeric variables have a value of 0 before assignment. String variables have a default value of "", which is also known as the empty string because it contains no characters. Boolean variables have a default value of false. 

```
var classTime uint32
var averageGrade float32
var teacherName string
var isPassFail bool

fmt.Println(classTime) // Prints 0
fmt.Println(averageGrade) // Prints 0
fmt.Println(teacherName) // Doesn't print anything
fmt.Println(isPassFail) // Prints false
```

### Inferring Type

There is a way to declare a variable without explicitly stating its type using the short declaration := operator. We might use the := operator if we know what value we want our variable to store when creating it.

```
nuclearMeltdownOccurring := true
radiumInGroundWater := 4.521
daysSinceLastWorkplaceCatastrophe := 0
externalMessage := "Everything is normal. Keep calm and carry on."
```

We can also perform the same inferring using : 

```
var nuclearMeltdownOccurring = true
var radiumInGroundWater = 4.521
var daysSinceLastWorkplaceCatastrophe = 0
var externalMessage = "Everything is normal. Keep calm and carry on."
```


### Multiple Variable Declaration

```
var part1, part2 string
quote, fact := "Bears, Beets, Battlestar Galactica", true
```



### The if Statement

```
alarmRinging := true
if alarmRinging {
  fmt.Println("Turn off the alarm!!") 
}

if (alarmRinging) {
  fmt.Println("Turn off the alarm!!") 
} else if otherCondition {
    fmt.Println("Do something else") 
}

isHungry := false
if isHungry {
  fmt.Println("Eat the cookie") 
} else {
  fmt.Println("Step away from the cookie...")
}

```

### Comparison Operators

```
"password1" == "password1" // Evaluates to true
"password1" == "badpass"   // Evaluates to false
123 != 12 // Evaluates to true
123 != 123 // Evaluates to false
100 < 200 // Evaluates to true
100.5 <= 100.5 // Evaluates to true
```

### Logical Operators (And, Or)

```
&&	And
||	Or
!	Not

if storeLights == "on" && doorsOpen {
  fmt.Println("You can enter the store!")
} 

```

### Functions

A simple function declaration goes by : 
```
func multiplier(x int32, y int32) int32 {
  return x * y
}
```


### Multiple Return Values

```
func GPA(midtermGrade float32, finalGrade float32) (string, float32) {
  averageGrade := (midtermGrade + finalGrade) / 2
  var gradeLetter string
  if averageGrade > 90 {
    gradeLetter = "A"
  } else if averageGrade > 80 {
    gradeLetter = "B"
  } else if averageGrade > 70 {
    gradeLetter = "C"
  } else if averageGrade > 60 {
    gradeLetter = "D"
  } else {
    gradeLetter = "F"
  }

  return gradeLetter, averageGrade 
}

myGrade, myAverage = GPA(89.4, 76.8)
```


### Deferring Resolution

We can delay a function call to the end of the current scope by using the defer keyword. defer tells Go to run a function, but at the end of the current function. This is useful for logging, file writing, and other utilities.

```
func deferTesting() {
	defer fmt.Println("Fikom a khouya chi f3ayl d nissa2")
	defer fmt.Println("Ntiya ra qtita Siamois")
	fmt.Println("Kiraaaaaa")
	fmt.Println("Chno bghitit nclipi rap")
}
```


### Addresses

To find a variable’s address we use the & operator followed by the variable name, like so :
```
x := "My very first address"
fmt.Println(&x) // Prints 0x414020
```


### Pointers

Pointers are variables that specifically store addresses.
```
var pointerForInt *int
minutes := 525600
pointerForInt = &minutes
fmt.Println(pointerForInt) // Prints 0xc000018038
fmt.Println(*pointerForInt) // Prints 525600

star := "Polaris"
var starAddress *string = &star
```

### Dereferencing
We know that addresses are where values are stored and pointers allow us to keep track of addresses. But what if we want the address to store a different value.

```
lyrics := "Moments so dear" 
pointerForStr := &lyrics
*pointerForStr = "Journeys to plan" 
fmt.Println(lyrics) // Prints: Journeys to plan
```


### For Loop

```
for number := 0; number < 5; number++ {
  fmt.Print(number, " ")
}

for count := 1; count <= 12; count=count+2 {
    fmt.Println(count)
}
```



### For as a While Loop
Most programming languages use the "while" keyword to declare indefinite loops. Go uses "for"

```
number := 0 // Initialize a variable to be used inside the loop
for number < 5 {
  fmt.Println(number)
  number++ // Update the variable being used
}
```



### Infinite Loops

An infinite loop is a loop where the condition can never be reached, causing the loop to run forever.
```
for {
  // Loop body logic
  // This repeats forever
}
// This is never reached
```


### Loop Keywords: Break and Continue

The "break" keyword allows the programmer to stop the loop at the current iteration. The "continue" keyword works similarly, allowing the loop to skip to the next iteration.


```
animals := []string{"Cat", "Dog", "Fish", "Turtle"}
for index := 0; index < len(animals); index++ {
  if animals[index] == "Dog" {
    fmt.Println("Found the perfect animal!")
    break // Stop searching the array
  }
}

jellybeans := []string{"green", "blue", "yellow", "red", "green", "yellow", "red"}
for index := 0; index < len(jellybeans); index++ {
  if jellybeans[index] == "green" {
    continue
  }
  fmt.Println("You ate the", jellybeans[index], "jellybean!")
}
```


### Arrays
When we declare a variable in Go, the compiler:
    - Finds space in memory for that variable
    - Associates the variable with a name

Declaring an array in Go requires that we provide the number of elements. Once declared, we cannot change this number without declaring a new array.


```
var arr1 [n]type
arr2 := [n]type{el1, el2, ..., eln}
var intArray []int  
floatArray := [4]float32{1.1,2.2,3.3,4.4}
zeroes := [3]float64{}
triangleSides := [3]int{15, 26, 30}
triangleAngles := [...]int{30, 60, 90} // the compiler auto determine the size of the array
```


### Slices
Slices are a data collection type similar to arrays, but they have the ability to change their size.
```
var intArray []int
stringSlice := []string{}
names := []string{"Kathryn", "Martin", "Sasha", "Steven"}
```

A slice is resizeable, so there is a difference between:
- Its length, the current number of elements it holds
- Its capacity, the number of elements it can hold before needing to resize itself.

```
slice := []string{"Fido", "Fifi", "FruFru"}
fmt.Println(slice, len(slice), cap(slice))
```


How to add an element to a slice :
```
books := []string{"Tom Sawyer", "Of Mice and Men"}
books = append(books, "Frankenstein")
books = append(books, "Dracula")
```




### Maps 

Go's implementation of hash tables.

To declare an empty map :
```
prices := make(map[string]float32)
highs := map[int]int{}
```



```
variable_name := map[key_data_type]value_data_type{}
map_name := map[key_data_type]value_data_type{
  key-1: value-1,
  key-2: value-2,
  key-N: value-N,
}
variable_name := map_name[key_value] 
```


Each map and array has a set amount of items that they contain. In Go, the range keyword can be used to work through these items one at a time within a loop. For example:
```
letters := []string{"A", "B", "C", "D"}
for index, value := range letters {
  fmt.Println("Index:", index, "Value:", value)
}
```


To check if a key is in a map we use a bool :
```
customer,status := customers["billy"]

if status {
  fmt.Println("we found the customer")
} else {
  fmt.Println("no such customer!")
}
```


To add an element to a map :
```
customers["simo"] = 345
```


To delete an key-value pair from a map :
```
delete(yourMap, keyValueToDelete)
delete(contacts, "Gary")
```





