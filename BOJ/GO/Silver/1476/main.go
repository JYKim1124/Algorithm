package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var E, S, M int
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var year = 1

	fmt.Fscanln(reader, &E, &S, &M)

	for {
		if (year-E)%15 == 0 && (year-S)%28 == 0 && (year-M)%19 == 0 {
			break
		}
		year++
	}

	fmt.Fprintln(writer, year)
}
