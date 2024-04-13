

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


