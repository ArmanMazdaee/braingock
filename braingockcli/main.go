package main

import (
	"context"
	"flag"
	"io"
	"log"
	"os"
	"os/signal"

	"github.com/ArmanMazdaee/braingock"
	"github.com/ArmanMazdaee/braingock/commands"
	"github.com/ArmanMazdaee/braingock/state"
)

var arraySize = flag.Uint("array-size", 0, "set the state array to a fixed size")
var leftExpandable = flag.Bool("left-expandable", false, "indecate if state can be expanded from the left")

func createState[T state.StateType]() state.State[T] {
	if *arraySize == 0 && *leftExpandable {
		log.Fatalln("both -array-size and -left-expandable options can not be set")
	}

	if *arraySize != 0 {
		return state.NewFixedSize[T](*arraySize)
	}
	return state.NewDynamicSize[T](*leftExpandable)
}

var inputPath = flag.String("input", "-", "file to be used for programs inputs")
var outputPath = flag.String("output", "-", "file to be used for programs outputs")
var errorPath = flag.String("error", "-", "file to be used for programs errors")

func createCommands[T state.StateType]() map[rune]commands.Command[T] {
	cmds := map[rune]commands.Command[T]{
		'>': commands.NewIncrementPointer[T](),
		'<': commands.NewDecrementPointer[T](),
		'+': commands.NewIncrementPointerValue[T](),
		'-': commands.NewDecrementPointerValue[T](),
		'[': commands.NewStartLoop[T](),
		']': commands.NewEndLoop[T](),
	}

	var errWriter io.Writer = os.Stderr
	if *errorPath != "-" {
		if f, err := createFile(*inputPath); err != nil {
			log.Fatalln("could not create err file:", err)
		} else {
			errWriter = f
		}
	}

	var inputReader io.Reader = os.Stdin
	var promptWriter io.Writer = os.Stdout
	if *inputPath != "-" {
		if f, err := openFile(*inputPath); err != nil {
			log.Fatalln("could not open input file:", err)
		} else {
			inputReader = f
			promptWriter = nil
		}
	}
	cmds[','] = commands.NewReadByte[T](inputReader, promptWriter, errWriter)

	var outputWriter io.Writer = os.Stdout
	if *outputPath != "-" {
		if f, err := createFile(*inputPath); err != nil {
			log.Fatalln("could not create output file:", err)
		} else {
			outputWriter = f
		}
	}
	cmds['.'] = commands.NewWriteByte[T](outputWriter, errWriter)

	return cmds
}

func run[T state.StateType]() {
	state := createState[T]()
	commands := createCommands[T]()
	inter := braingock.NewInterpreter(state, commands)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	errCh := orchestrateComputation(ctx, inter, flag.Args())

	interruptCh := make(chan os.Signal, 1)
	signal.Notify(interruptCh, os.Interrupt)

	for {
		select {
		case err := <-errCh:
			if err == nil {
				log.Println("done")
				return
			}
			log.Println(err)
		case <-interruptCh:
			if err := ctx.Err(); err == context.Canceled {
				log.Fatalln("force shutdown")
			}
			cancel()
			log.Println("graceful shutdown...")
		}
	}
}

var cellSize = flag.Uint("cell-size", 8, "the word size of each state call. should be one of: 8, 16, 32, 64")

func main() {
	flag.Parse()
	switch *cellSize {
	case 8:
		run[uint8]()
	case 16:
		run[uint16]()
	case 32:
		run[uint32]()
	case 64:
		run[uint64]()
	}
}
