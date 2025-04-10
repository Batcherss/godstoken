package util

import "fmt"

func Debug(enabled bool, args ...interface{}) {
	if enabled {
		fmt.Println(append([]interface{}{"[DEBUG]"}, args...)...)
	}
}
