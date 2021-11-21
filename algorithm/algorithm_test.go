package algorithm

import (
	"math"
	"testing"

	"github.com/stretchr/testify/suite"
)

type AlgorithmTest struct {
	suite.Suite
}

func (t *AlgorithmTest) TestReverseVertical() {
	t.Equal(
		[][]int{},
		reverseVertical(0, 0, [][]int{}),
	)
	t.Equal(
		[][]int{{1, 2}, {3, 4}},
		reverseVertical(2, 2, [][]int{{3, 4}, {1, 2}}),
	)
}

func (t *AlgorithmTest) TestReverseVerticalBool() {
	t.Equal(
		[][]bool{},
		reverseVerticalBool(0, 0, [][]bool{}),
	)
	t.Equal(
		[][]bool{{false, true}, {true, false}},
		reverseVerticalBool(2, 2, [][]bool{{true, false}, {false, true}}),
	)
}

func (t *AlgorithmTest) TestReverseHorizontal() {
	t.Equal(
		[][]int{},
		reverseHorizontal(0, 0, [][]int{}),
	)
	t.Equal(
		[][]int{{1, 2}, {3, 4}},
		reverseHorizontal(2, 2, [][]int{{2, 1}, {4, 3}}),
	)
}

func (t *AlgorithmTest) TestReverseHorizontalBool() {
	t.Equal(
		[][]bool{},
		reverseHorizontalBool(0, 0, [][]bool{}),
	)
	t.Equal(
		[][]bool{{true, true}, {true, false}},
		reverseHorizontalBool(2, 2, [][]bool{{true, true}, {false, true}}),
	)
}

func (t *AlgorithmTest) TestReverse() {
	t.Equal(
		[][]int{},
		reverse(0, 0, [][]int{}),
	)
	t.Equal(
		[][]int{{1, 2}, {3, 4}},
		reverse(2, 2, [][]int{{4, 3}, {2, 1}}),
	)
}

func (t *AlgorithmTest) TestReverseBool() {
	t.Equal(
		[][]bool{},
		reverseBool(0, 0, [][]bool{}),
	)
	t.Equal(
		[][]bool{{true, true}, {true, false}},
		reverseBool(2, 2, [][]bool{{false, true}, {true, true}}),
	)
}

func (t *AlgorithmTest) TestMin() {
	t.Equal(math.MaxInt32, min())
	t.Equal(1, min(1, 2, 3))
	t.Equal(2, min(7, 12, 2))
}

func (t *AlgorithmTest) TestCalculateLeftTopNearest() {
	t.Equal(
		[][]int{},
		calculateLeftTopNearest(0, 0, [][]bool{}),
	)
	t.Equal(
		[][]int{{0, 1}, {0, 1}},
		calculateLeftTopNearest(2, 2, [][]bool{{true, false}, {true, false}}),
	)
	t.Equal(
		[][]int{{math.MaxInt32, math.MaxInt32, 0}, {0, 1, 0}, {1, 2, 0}},
		calculateLeftTopNearest(3, 3, [][]bool{{false, false, true}, {true, false, true}, {false, false, true}}),
	)
}

func (t *AlgorithmTest) TestSolve() {
	t.Equal(
		[][]int{},
		Solve(0, 0, [][]bool{}),
	)
	t.Equal(
		[][]int{{}, {}, {}},
		Solve(3, 0, [][]bool{{}, {}, {}}),
	)
	t.Equal(
		[][]int{
			{0, 1, 1},
			{0, 0, 0},
			{1, 1, 0},
		},
		Solve(3, 3, [][]bool{
			{true, false, false},
			{true, true, true},
			{false, false, true}},
		),
	)
	t.Equal(
		[][]int{
			{3, 2, 1, 0},
			{2, 1, 0, 0},
			{1, 0, 0, 1},
		},
		Solve(3, 4, [][]bool{
			{false, false, false, true},
			{false, false, true, true},
			{false, true, true, false}},
		),
	)
}

func TestAlgorithm(t *testing.T) {
	suite.Run(t, new(AlgorithmTest))
}
