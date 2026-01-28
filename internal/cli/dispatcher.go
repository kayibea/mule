package cli

import (
	"fmt"
	"os"
)

func Run(args []string) int {
	if len(args) < 2 {
		printHelp(args[0])
		os.Exit(1)
	}

	cmd := args[1]
	cmdArgs := args[2:]

	switch cmd {
	case "add":
		return runAdd(cmdArgs)
	case "list":
		return runList(cmdArgs)
	case "copy":
		return runCopy(cmdArgs)
	case "move":
		return runMove(cmdArgs)
	case "prune":
		return runPrune(cmdArgs)
	case "help":
		printHelp(args[0])
		os.Exit(0)
	default:
		fmt.Printf("Invalid command '%s'; type \"help\" for a list\n", cmd)
		os.Exit(1)
	}

	return 0
}

func printHelp(prog string) {
	fmt.Println("Usage:")
	fmt.Printf("  %s <command> [options] [files...]\n", prog)
	fmt.Printf("  %s <command> --help\n\n", prog)

	fmt.Println("Commands:")
	fmt.Println("  add      Add files to the clipboard")
	fmt.Println("  list     List the clipboard content")
	fmt.Println("  copy     Copy mule files into current directory")
	fmt.Println("  move     Move mule files into current directory")
	fmt.Println("  prune    Clear the clipboard")
	fmt.Println("  help     Show this help")
}
