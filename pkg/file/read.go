package file

import (
	"bufio"
	"fmt"
	"os"

	"github.com/batcherss/godstoken/pkg/token"
	"github.com/batcherss/godstoken/pkg/util"
)

func ProcessFile(path string, debug bool, collectInfo bool) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("[-] Failed to open file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		util.Debug(debug, "[+] Checking line:", line)
		token.CheckToken(line, debug, collectInfo)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("[-] Error reading file:", err)
	}
}
