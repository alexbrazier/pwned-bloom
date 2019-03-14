package main

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"os"
	"strings"

	"github.com/willf/bloom"
)

const bloomFile = "out.txt"

func main() {
	filter := bloom.New(1, 1)
	fmt.Printf("Loading %s\n", bloomFile)
	f, _ := os.Open(bloomFile)
	defer f.Close()
	r := bufio.NewReader(f)
	filter.ReadFrom(r)
	scanner := bufio.NewReader(os.Stdin)
	fmt.Println("Enter password to test - Ctrl/Cmd + C to exit")
	for {
		fmt.Print("password: ")
		in, _ := scanner.ReadString('\n')
		password := strings.TrimSuffix(in, "\n")
		hash := strings.ToUpper(fmt.Sprintf("%x", sha1.Sum([]byte(password))))
		fmt.Printf("%s:%s\n", password, hash)
		found := filter.Test([]byte(hash))
		fmt.Printf("Password found: %t\n", found)
	}
}
