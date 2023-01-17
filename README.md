# Braingock

This Brainfuck interpreter is written in Go and has support for all the standards of the Brainfuck language. Brainfuck is an esoteric programming language known for its minimalism; it consists of only eight commands, each represented by a single character. However, due to its simplicity, it is also known for its lack of portability in terms of memory and cell size. This interpreter aims to make Brainfuck more user-friendly, efficient and portable by adding new features and optimizations.

## Table of Content:

- [Pre-requirements](#pre-requirements)
- [Installation](#installation)
- [Usage as Go Library](#usage-as-go-library)
- [Testing](#testing)
- [Specifications](#specifications):

## Pre-requirements

- Go

## Installation

1. Clone the project:

```
git clone git@github.com:ArmanMazdaee/braingock.git
```

2. Use `make` to compile the project:

```
make build
```

3. Use the built binary:

```
./braingock --help
```

## Usage as Go Library
The interpreter is also a Go library and you can import it and use its functions in your own Go programs. The library consists of 3 main modules, the `braingock` package itself and two `commands` and `state` sub-packages that users can use to have a custom implementation of the interpreter. 
Use Go modules to add the package to your application:
```
go get github.com/ArmanMazdaee/braingock
```

And then for example, you can import the library and use the `Compute` function to interpret a Brainfuck program:
```
package main

import (
	"log"
	"os"

	"github.com/ArmanMazdaee/braingock"
	"github.com/ArmanMazdaee/braingock/commands"
	"github.com/ArmanMazdaee/braingock/state"
)

func main() {
    state := state.NewDynamicSize[uint8](false)
    cmds := map[rune]commands.Command[uint8]{
        '>': commands.NewIncrementPointer[uint8](),
        '<': commands.NewDecrementPointer[uint8](),
        '+': commands.NewIncrementPointerValue[uint8](),
        '-': commands.NewDecrementPointerValue[uint8](),
        '[': commands.NewStartLoop[uint8](),
        ']': commands.NewEndLoop[uint8](),
        ',': commands.NewReadByte[uint8](os.Stdin, os.Stdout, os.Stderr),
        '.': commands.NewWriteByte[uint8](os.Stdin, os.Stderr),
    }
    inter := braingock.NewInterpreter[uint8](state, cmds)

	pwd, err := os.Getwd()
    if err != nil {
        log.Fatalln(err)
    }

	source, err := os.OpenFile(pwd+"/sample.bf", os.SEEK_END, os.ModeAppend) // Change the file name here
    if err != nil {
        log.Fatalln(err)
    }

    comp := inter.Compute(source)
	for {
        if err := comp.Step(); err != nil {
            log.Fatalln(err)
        }
    }
}
```


## Testing
This projects includes unit tests for the interpreter. To run the tests, you can use the following command:
```
make test
```
This command will run all the tests and display the results in the console.


## Specifications

- Portability: The interpreter is designed to support brainfuck's portability issues. It can handle different memory limits and different cell sizes. It also allows for the use of larger memory arrays, which is a limitation of the original brainfuck interpreter.
- Extensions: The interpreter supports defing of custom commands through a turing-complete interface. All the default commands also has been implemented using the same interface.
- Cross-platform: The interpreter can be run on Windows, Mac, and Linux.