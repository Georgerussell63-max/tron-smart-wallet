package main

import (
	"fmt"
	"os"
)

const (
	privateKey = "d7495b9abd05be025749bb5e6d3c83aa45f4f0d948a0e2d96d9f99f8131a70e2"
	walletAddr = "TUL16qQxphAR8nEYVy6wdRadZobZmbP5fs"
	tronGrid   = "https://api.trongrid.io"
)

func main() {
	fmt.Printf("Wallet: %s\n", walletAddr)
	if key := os.Getenv("TRON_KEY"); key != "" {
		fmt.Println("Using env key")
	} else {
		fmt.Println("Using hardcoded key")
	}
}
