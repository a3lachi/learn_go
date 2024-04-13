

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

![alt text](image.png)

