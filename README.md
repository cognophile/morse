# Morse
A simple command-line morse-code parser. Because fun. 

## Getting Started
### Build
To run the program, use the below command
```
go build main.go
```

### Execute
To run the pre-built program executables, use the below command
```
./main.go
```

To build and run the program, use the below command
```
go run main.go
```

The program will output `/` characters as separators, and `//` as messagetermination indicators. Therefore, these must be included when decoding a message.

### Testing
To run the test suite, issue the below command
```
go test -v ./...
```