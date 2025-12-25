// Package sortutil provides helpers for sorting strings read from a file.
package sortutil

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

// Sort defines options for sorting strings.
type Sort struct {
	strs                []string
	compare             func(a, b string) int
	print               func()
	col                 int
	unique              bool
	ignoreLeadingBlanks bool
	isSorted            bool
}

// New returns a new Sort with default settings.
func New() *Sort {
	s := &Sort{
		strs:                make([]string, 0, 100),
		compare:             compareStrs,
		col:                 1,
		unique:              false,
		ignoreLeadingBlanks: false,
		isSorted:            false,
	}
	s.print = s.printDefault

	return s
}

// Append adds a string to the sorter.
func (s *Sort) Append(str string) {
	s.strs = append(s.strs, str)
}

// AsNums enables numeric comparison of strings.
func (s *Sort) AsNums() {
	s.compare = compareNums
}

// AsMonths enables month-based string comparison.
func (s *Sort) AsMonths() {
	s.compare = compareMonths
}

// AsHumanNums enables human-readable numeric comparison of strings.
func (s *Sort) AsHumanNums() {
	s.compare = compareHumanNums
}

// Unique enables unique string sorting.
func (s *Sort) Unique() {
	s.unique = true
}

// IgnoreLeadingBlanks ignores leading blanks when comparing strings.
func (s *Sort) IgnoreLeadingBlanks() {
	s.ignoreLeadingBlanks = true
}

// IsSorted enables sorted-check mode.
func (s *Sort) IsSorted() {
	s.print = s.printIsSorted
	s.isSorted = true
}

// Reverse enables reverse sorting mode.
func (s *Sort) Reverse() {
	s.print = s.printReverse
}

// Print prints the sorting results.
func (s *Sort) Print() {
	s.print()
}

// Sort sorts strings by the specified column.
func (s *Sort) Sort(col int) {
	s.col = col - 1
	if s.isSorted {
		s.isSorted = slices.IsSortedFunc(s.strs, s.cmp)
	} else {
		slices.SortFunc(s.strs, s.cmp)
	}
	if s.unique {
		s.strs = slices.Compact(s.strs)
	}
}

func (s *Sort) cmp(a, b string) int {
	if s.ignoreLeadingBlanks {
		a = strings.TrimSpace(a)
		b = strings.TrimSpace(b)
	}
	if s.col > 0 {
		colsA := strings.Split(a, "\t")
		colsB := strings.Split(b, "\t")

		if len(colsA) > s.col && len(colsB) > s.col {
			return s.compare(colsA[s.col], colsB[s.col])
		}
	} else {
		return s.compare(a, b)
	}

	return 0
}

func (s *Sort) printDefault() {
	for i, str := range s.strs {
		fmt.Printf("%v: %v\n", i+1, str)
	}
}

func (s *Sort) printReverse() {
	for i := len(s.strs) - 1; i >= 0; i-- {
		fmt.Printf("%v: %v\n", i+1, s.strs[i])
	}
}

func (s *Sort) printIsSorted() {
	fmt.Println(s.isSorted)
}

func compareStrs(a, b string) int {
	return cmp.Compare(a, b)
}

func compareNums(a, b string) int {
	intA, err := strconv.Atoi(a)
	if err != nil {
		return 0
	}
	intB, err := strconv.Atoi(b)
	if err != nil {
		return 0
	}

	return cmp.Compare(intA, intB)
}

func compareMonths(a, b string) int {
	a = strings.ToLower(a)
	b = strings.ToLower(b)
	intA := 0
	intB := 0

	for _, month := range months {
		if strings.Contains(a, month.name) {
			intA = int(month.value)
		}
		if strings.Contains(b, month.name) {
			intB = int(month.value)
		}
	}

	return cmp.Compare(intA, intB)
}

func compareHumanNums(a, b string) int {
	a = strings.ToLower(a)
	b = strings.ToLower(b)

	numA := 0
	intA := 0
	indexA := -1
	numB := 0
	intB := 0
	indexB := -1

	for _, num := range humanNums {
		indexA = strings.Index(a, num.name)
		if indexA != -1 {
			numA = int(num.value)
			intA, _ = strconv.Atoi(a[0:indexA])
		}

		indexB = strings.Index(b, num.name)
		if indexB != -1 {
			numB = int(num.value)
			intB, _ = strconv.Atoi(b[0:indexB])
		}
	}

	if numA == numB {
		return cmp.Compare(intA, intB)
	}

	return cmp.Compare(numA, numB)
}
