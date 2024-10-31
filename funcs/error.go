package lemIn

import "fmt"

func Error(msg string) {
	fmt.Println("\033[31mERROR: invalid data format,", msg, "\033[0m")
}
