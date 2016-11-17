package main

import (
	"flag"
	"fmt"
	"path/filepath"
)

var (
	pattern = flag.String("pattern", "", "Healp!")
)

func main() {
	flag.Parse()
	if *pattern == "" {
		flag.PrintDefaults()
	}
	results, err := filepath.Glob(*pattern)
	if err != nil {
		fmt.Printf("%v\n", results)
	}
	for _, result := range results {
		fmt.Printf("%v\n", result)
	}
}
