package drago

import (
	"drago/internal/injector"
	"fmt"
	"os"
)

func help() {
	fmt.Println("Usage: drago [command] " +
		"")
}

func version() {
	fmt.Println("go drago version 1.0.0")
}

func inject() {
	if len(os.Args) < 2 {
		fmt.Printf("too few arguments\n")
	}
	paths := os.Args[2:]
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
		fmt.Printf("Unknown command: %s\n", command)
	}

}
