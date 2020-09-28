// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	lineFiles := make(map[string]map[string]bool)
	if len(files) == 0 {
		countLines(os.Stdin, counts, lineFiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, lineFiles)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			for file := range lineFiles[line] {
				fmt.Println(file)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int, lineFiles map[string]map[string]bool) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		filename := f.Name()
		counts[line]++
		lineFiles[line][filename] = true
	}
	// NOTE: ignoring potential errors from input.Err()
}
