package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/suite"
)

type AlgorithmTest struct {
	suite.Suite
}

func (t *AlgorithmTest) TestReadInputZeroCases() {
	reader := bytes.NewBuffer([]byte("0"))

	result := readInput(reader)

	t.Equal(result, []input{})
}

func (t *AlgorithmTest) TestReadInput() {
	reader := bytes.NewBuffer([]byte(`2
	2 2
	01
	10
	3 1
	0
	1
	0`))

	result := readInput(reader)

	t.Equal(result, []input{
		{2, 2, [][]bool{{false, true}, {true, false}}},
		{3, 1, [][]bool{{false}, {true}, {false}}},
	})
}

func (t *AlgorithmTest) TestOutput() {
	writer := bytes.NewBuffer([]byte{})
	expected := `1 2 3 4
5 6 7 8
`

	output(writer, [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}})

	t.Equal(expected, writer.String())
}

func (t *AlgorithmTest) TestOutputEmptyArray() {
	writer := bytes.NewBuffer([]byte{})
	expected := ``

	output(writer, [][]int{})

	t.Equal(expected, writer.String())
}

func (t *AlgorithmTest) TestOutputMultipleEmptyLines() {
	writer := bytes.NewBuffer([]byte{})
	expected := `


`

	output(writer, [][]int{{}, {}, {}})

	t.Equal(expected, writer.String())
}

func TestAlgorithm(t *testing.T) {
	suite.Run(t, new(AlgorithmTest))
}
