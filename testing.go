package aoc2023

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"testing"
)

type SolverTest struct {
	day           int
	version       string
	subversion    string
	solver        Solver
	expectPartOne Result
	expectPartTwo Result
}

func (st SolverTest) Name() string {
	name := fmt.Sprintf("day=%d;version=%s", st.day, st.version)
	if st.subversion != "" {
		name = fmt.Sprintf("%s.%s", name, st.subversion)
	}
	return name
}
func (st SolverTest) scanInput() *bufio.Scanner {
	path := fmt.Sprintf("testdata/%s/%d", st.version, st.day)
	if st.subversion != "" {
		path = fmt.Sprintf("%s.%s", path, st.subversion)
	}
	return bufio.NewScanner(Must(os.Open(path)))
}
func (st SolverTest) Go(t *testing.T) {
	st.solver.Load(st.scanInput())
	st.testSolver(t, st.expectPartOne, st.solver.PartOne)
	st.testSolver(t, st.expectPartTwo, st.solver.PartTwo)
}

func (st SolverTest) testSolver(t *testing.T, expected Result, solve func() Result) {
	if expected == nil {
		return
	}
	received := solve()
	switch expected.(type) {
	case NumberResult:
		expectedValue, _ := expected.Number()
		receivedValue, err := received.Number()
		if err != nil {
			t.Errorf("failed: %v", err)
		} else if expectedValue != receivedValue {
			t.Errorf("expected %d but received %d", expectedValue, receivedValue)
		}
	case StringResult:
		expectedValue, _ := expected.String()
		receivedValue, err := received.String()
		if err != nil {
			t.Errorf("failed: %v", err)
		} else if expectedValue != receivedValue {
			t.Errorf("expected %q but received %q", expectedValue, receivedValue)
		}
	case ErrorResult:
		expectedValue, _ := expected.Error()
		receivedValue, err := received.Error()
		if err != nil {
			t.Errorf("failed: %v", err)
		} else if !errors.Is(expectedValue, receivedValue) {
			t.Errorf("expected %v but received %v", expectedValue, receivedValue)
		}
	}
}
