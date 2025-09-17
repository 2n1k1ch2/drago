package drago

import (
	"drago/internal/injector"
	"fmt"
	"os"
)

func help() {
	fmt.Println("Usage: drago [command]")
}

func version() {
	fmt.Println("go drago version 1.0.0")
}

func inject() {
	paths := os.Args[1:]
	injector.Build(paths)
}
func main() {
	command := os.Args[1]
	switch command {
	case "help":
		help()
	case "version":
		version()
	case "inject":
		inject()
	default:
		fmt.Printf("Unkown command: %s\n", command)
	}

}
