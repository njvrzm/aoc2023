package aoc2023

import (
	"bufio"
	"errors"
)

type Solver interface {
	PartOne() Result
	PartTwo() Result
	Load(scanner *bufio.Scanner)
}

var WrongResultType = errors.New("wrong value type requested from result")

type Result interface {
	Number() (int, error)
	String() (string, error)
	Error() (error, error)
}

type NumberResult struct {
	value int
}

func (n NumberResult) Number() (int, error) {
	return n.value, nil
}

func (n NumberResult) String() (string, error) {
	return "", WrongResultType
}

func (n NumberResult) Error() (error, error) {
	return nil, WrongResultType
}

type StringResult struct {
	value string
}

func (s StringResult) Number() (int, error) {
	return 0, WrongResultType
}

func (s StringResult) String() (string, error) {
	return s.value, nil
}

func (s StringResult) Error() (error, error) {
	return nil, WrongResultType
}

type ErrorResult struct {
	value error
}

func (e ErrorResult) Number() (int, error) {
	return 0, WrongResultType
}

func (e ErrorResult) String() (string, error) {
	return "", WrongResultType
}

func (e ErrorResult) Error() (error, error) {
	return e.value, nil
}

var NotImplemented = ErrorResult{errors.New("not implemented")}

type SkipResult struct{}

func (s SkipResult) Number() (int, error)    { return 0, nil }
func (s SkipResult) String() (string, error) { return "", nil }
func (s SkipResult) Error() (error, error)   { return nil, nil }
