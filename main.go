package main

import (
	"fmt"
	"github.com/batcherss/godstoken/pkg/token"
	"github.com/batcherss/godstoken/pkg/file"
	"flag"
)

func brainrootamobea() {
	banner := []string{
		"                   .___        __          __                  ",
		"   ____   ____   __| _/_______/  |_  ____ |  | __ ____   ____  ",
		"  / ___\\ /  _ \\ / __ |/  ___/\\   __\\/  _ \\|  |/ // __ \\ /    \\ ",
		" / /_/  >  <_> ) /_/ |\\___ \\  |  | (  <_> )    <\\  ___/|   |  \\",
		" \\___  / \\____/\\____ /____  > |__|  \\____/|__|_ \\\\___  >___|  /",
		"/_____/             \\/    \\/                   \\/    \\/     \\/ ",
		"   -god?.discord?.token? piece of sh#t name but ok.\n",
	}

	for _, line := range banner {
		fmt.Println(line)
	}
}


func printHelp() {
	brainrootamobea()
	fmt.Println("Usage: godstoken [options]")
	fmt.Println("-t=\"<token>\"     : Token to check (single token check)")
	fmt.Println("-f=\"<file_path>\" : Path to file with tokens (one per line)")
	fmt.Println("-d               : Gather user info (username, email, etc.)")
	fmt.Println("-debug           : Enable debug mode (shows detailed logs)")
	fmt.Println("Conflict: -t and -f cannot be used together")
}

func main() {
	tokenPtr := flag.String("t", "", "Token to check")
	filePtr := flag.String("f", "", "File path with tokens to check")
	debugPtr := flag.Bool("debug", false, "Enable debug mode")
	gatherInfoPtr := flag.Bool("d", false, "Gather user info")
	flag.Parse()

	brainrootamobea()

	if *tokenPtr != "" && *filePtr != "" {
		fmt.Println("[-] Conflict: -t and -f cannot be used together.")
		printHelp()
		return
	}

	if *tokenPtr == "" && *filePtr == "" {
		fmt.Println("[-] Error: You must specify either -t or -f.")
		printHelp()
		return
	}

	if *tokenPtr != "" {
		token.CheckToken(*tokenPtr, *debugPtr, *gatherInfoPtr)
	}

	if *filePtr != "" {
		file.ProcessFile(*filePtr, *debugPtr, *gatherInfoPtr)
	}
}
