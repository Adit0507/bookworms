package main

import (
	"fmt"
	"os"
)

func main() {
	bookworms, err := loadBookworms("testdata/bookworms.json")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load bookwroms:")
		os.Exit(1)
	}

	commonBooks := findCommonBooks(bookworms)
	fmt.Println("Here are the books in common:")
	displayBooks(commonBooks)
}