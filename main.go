package main

import (
	"fmt"
	"io"
	"noveo_test/algorithm"
	"os"
)

type input struct {
	n, m int
	data [][]bool
}

func readInput(reader io.Reader) []input {
	cases := 0
	fmt.Fscan(reader, &cases)
	result := make([]input, cases)
	for i := 0; i < cases; i++ {
		var n, m int
		fmt.Fscan(reader, &n, &m)
		result[i] = input{
			n:    n,
			m:    m,
			data: readBitmap(reader, n, m),
		}
	}
	return result
}

func readBitmap(reader io.Reader, n, m int) [][]bool {
	result := make([][]bool, n)
	for i := 0; i < n; i++ {
		result[i] = make([]bool, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &result[i][j])
		}
	}
	return result
}

func output(writer io.Writer, data [][]int) {
	for _, row := range data {
		for i, val := range row {
			fmt.Fprint(writer, val)
			if i != len(row)-1 {
				fmt.Fprint(writer, " ")
			}
		}
		fmt.Fprintln(writer)
	}
}

func main() {
	input := readInput(os.Stdin)
	for _, v := range input {
		output(os.Stdout, algorithm.Solve(v.n, v.m, v.data))
	}
}
