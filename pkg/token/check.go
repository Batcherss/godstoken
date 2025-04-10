package token

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"github.com/batcherss/godstoken/pkg/util"
)

type UserInfo struct {
	Username string `json:"username"`
	ID       string `json:"id"`
}

func CheckToken(token string, debug bool, collectInfo bool) {
	util.Debug(debug, "[+] Checking token", token)

	req, err := http.NewRequest("GET", "https://discord.com/api/v10/users/@me", nil)
	if err != nil {
		util.Debug(debug, "[-] Failed to create request:", err)
		fmt.Printf("[-] Token \"%s\" invalid.\n", token)
		return
	}

	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")

	util.Debug(debug, "[+] Connecting to Discord API...")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		util.Debug(debug, "[-] Request error:", err)
		fmt.Printf("[-] Token \"%s\" invalid.\n", token)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		util.Debug(debug, "[-] Invalid response status:", resp.Status)
		fmt.Printf("[-] Token \"%s\" invalid.\n", token)
		return
	}

	util.Debug(debug, "[+] Token valid:", token)
	fmt.Printf("[+] Token \"%s\" is valid.\n", token)

	file, err := os.OpenFile("active_tokens.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		util.Debug(debug, "[-] Failed to open file:", err)
		fmt.Printf("[-] Failed to open active_tokens.txt.\n")
		return
	}
	defer file.Close()

	_, err = file.WriteString(token + "\n")
	if err != nil {
		util.Debug(debug, "[-] Failed to write to file:", err)
		fmt.Printf("[-] Failed to write token to active_tokens.txt.\n")
		return
	}

	if collectInfo {
		body, _ := ioutil.ReadAll(resp.Body)
		var info UserInfo
		json.Unmarshal(body, &info)

		fmt.Println("\n[ INFO ]")
		fmt.Println("Token    -", token)
		fmt.Println("Username -", info.Username)
		fmt.Println("ID       -", info.ID)
	}
}
