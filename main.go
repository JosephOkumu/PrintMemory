package main

import (
	"fmt"
	"os"
)

func PrintMemory(arr [10]byte) {
	for i := 0; i < 10; i++ {
		if i > 0 && i%4 == 0 {
			os.Stdout.WriteString("\n")
		}
		if i < 9 {
			os.Stdout.WriteString(fmt.Sprintf("%02x ", arr[i]))
		} else {
			os.Stdout.WriteString(fmt.Sprintf("%02x", arr[i]))
		}
	}
	os.Stdout.WriteString("\n")

	for _, b := range arr {
		if b >= 32 && b <= 126 {
			os.Stdout.WriteString(string(b))
		} else {
			os.Stdout.WriteString(".")
		}
	}
	os.Stdout.WriteString("\n")
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
