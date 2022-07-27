package main

import (
	"fmt"
	"os"

	"github.com/manoj-gupta/go-bloom/bloomfilter"
)

func main() {
	fang := []string{
		"www.facebook.com",
		"www.amazon.com",
		"www.google.com",
		"www.netflix.com",
	}
	notFang := "not-fang"

	// create bloomfilter instance
	bf := bloomfilter.New(100, bloomfilter.DefaultHashFunctions)

	// Add `fang` to bloomfilter
	for _, company := range fang {
		bf.Add([]byte(company))
	}

	// Check if `fang` exist in bloomfilter
	for _, company := range fang {
		found := bf.Search([]byte(company))
		if !found {
			fmt.Printf("%s NOT found\n", company)
			os.Exit(1)
		}
	}

	// check for something which is not `fang`
	found := bf.Search([]byte(notFang))
	if found {
		fmt.Printf("%s found .. should not exist\n", notFang)
		os.Exit(1)
	}

	fmt.Println("It Works!!")
}
