package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"

	"github.com/willf/bloom"
)

const inFile = "pwned-passwords-sha1-ordered-by-hash-v4.txt"
const outFile = "out.txt"

func main() {
	n := float64(551509767)                                   // number of passwords
	p := 0.01                                                 // 1% accepted false positive rate
	mFloat := -n * math.Log2(p) / (math.Pow(math.Log2(2), 2)) // number of bits needed for bloom algorithm
	k := uint(math.Ceil(mFloat / n * math.Log2(2)))           // number of hash functions to apply
	m := uint(math.Ceil(mFloat))
	fmt.Printf("For %.0f passwords, and a false positive rate of %.3f, using %d bits and %d hash functions\n", n, p, m, k)
	generateBloom(inFile, m, k)
}

func generateBloom(filename string, m, k uint) {
	file, err := os.Open(filename)
	if err != nil {
		println(err)
	}
	defer file.Close()
	filter := bloom.New(m, k)

	scanner := bufio.NewScanner(file)
	fmt.Println("Running")

	count := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")
		hash := line[0]
		filter.Add([]byte(hash))
		count++
		if count%1000000 == 0 {
			fmt.Printf("%d - %s\n", count, hash)
		}
	}
	fmt.Println("Done!")
	f, _ := os.Create(outFile)
	defer f.Close()
	f.Sync()
	w := bufio.NewWriter(f)
	fmt.Println(filter.WriteTo(w))
	w.Flush()
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
