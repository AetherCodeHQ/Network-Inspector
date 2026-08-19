package main

import (
	"fmt"
	"os"
)

// network_inspector - Inspect network traffic
func network_inspector(path string) {
	fmt.Println("========================================")
	fmt.Println("  Network-Inspector")
	fmt.Println("  Inspect network traffic")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	network_inspector(path)
}
