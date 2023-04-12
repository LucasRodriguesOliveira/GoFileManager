package log

import "fmt"

func Info(msg string) {
	fmt.Printf("[INFO]: %v\n", msg)
}
