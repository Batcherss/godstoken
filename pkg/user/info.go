package user

import (
	"fmt"
)

type UserInfo struct {
	Username     string `json:"username"`
	Discriminator string `json:"discriminator"`
	Email        string `json:"email"`
}

func PrintUserInfo(token string, userInfo UserInfo) {
	fmt.Println("[ INFO ]")
	fmt.Println("Token -", token)
	fmt.Println("Username -", userInfo.Username+"#"+userInfo.Discriminator)
	fmt.Println("Email -", userInfo.Email)
}
