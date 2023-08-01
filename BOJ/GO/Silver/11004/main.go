package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var N, K int
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N, &K)

	// N을 입력받고 배열 선언해야 함
	var arr []int = make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &arr[i])
	}
	sort.Ints(arr)

	fmt.Fprintln(writer, arr[K-1])

}
