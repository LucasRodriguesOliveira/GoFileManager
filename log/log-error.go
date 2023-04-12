package log

import "fmt"

func Error(msg string) {
	fmt.Printf("[ERR]: %v\n", msg)
}
