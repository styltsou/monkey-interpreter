package repl

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/styltsou/monkey-interpreter/lexer"
	"github.com/styltsou/monkey-interpreter/token"
)

const PROMPT = ">>"

func clearConsole() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "darwin", "linux":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func Start() {
	// Register a signal handler for interrupt and termination signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// Start a goroutine to wait for the signal and exit
	go func() {
		<-signalChan
		os.Exit(0)
	}()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(PROMPT)

		if scanner.Scan() {
			line := scanner.Text()

			switch line {
			case "clear":
				clearConsole()
			case "exit":
				signalChan <- syscall.SIGINT
				break
			default:
				lexer := lexer.New(line)
				// Here gather all the tokens in an array
				// pretty print the tokens
				for tok := lexer.NextToken(); tok.Type != token.EOF; tok = lexer.NextToken() {
					fmt.Printf("%+v\n", tok)

				}
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error:", err)
			break
		}
	}
}
