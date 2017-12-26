package main

import (
	"data"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type arguments struct {
	isQuote    bool
	path, font string
	color      int8
}

// If arguments are provided, runs episodically
func validate(args []string) (argumentList arguments, err error) {
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--quotes":
			argumentList.isQuote = true
		case "-S", "--src":
			argumentList.path = args[i+1]
			i = i + 1
		case "-F", "--font":
			argumentList.font = args[i+1]
			i = i + 1
		case "-C", "--font-color":
			colorReceived, _ := strconv.Atoi(args[i+1])
			argumentList.color = int8(colorReceived)
			i = i + 1
		default:
			return argumentList, errors.New("Invalide argument " + args[i])
		}
	}

	return argumentList, nil
}

// Reads the configuration file. Daemonizes the program
func readTheConfigurations() (argumentList arguments, err error) {
	return
}

/*
	Arguments
	--quotes (for quotes, default is word of the day)
	-S --src /path/to/image (default=random color)
	-F --font /path/to/font (default=)
	-C --font-color RRGGBB

*/

func main() {
	daemon := false

	argumentsProvided := os.Args[1:]
	if len(argumentsProvided) == 0 {
		daemon = true
	}

	var err error
	var args arguments

	if !daemon {
		args, err = validate(argumentsProvided)
	} else {
		args, err = readTheConfigurations()
	}

	if err != nil {
		panic(err)
	}

	var text *http.Response
	var errr error

	if args.isQuote {
		text, errr = data.GetQuote()
	} else {
		text, errr = data.GetWord()
	}

	if errr != nil {
		panic(err)
	}

	fmt.Print(text)
}
