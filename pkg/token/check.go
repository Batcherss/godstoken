package token

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
