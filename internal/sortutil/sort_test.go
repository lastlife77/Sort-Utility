package sortutil

import (
	"slices"
	"testing"
)

func TestSortDefault(t *testing.T) {
	data := []string{"dog", "cat", " rat", "dog", "cow", "rat"}
	exp := []string{" rat", "cat", "cow", "dog", "dog", "rat"}

	s := New()
	for _, str := range data {
		s.Append(str)
	}
	s.Sort(1)
	act := s.strs
	if !slices.Equal(exp, act) {
		t.Fatalf("\nExpected:\n %v\nActual:\n %v", exp, act)
	}
}

func TestSortNumbers(t *testing.T) {
	data := []string{"45", "2", "556", "7"}
	exp := []string{"2", "7", "45", "556"}

	s := New()
	for _, str := range data {
		s.Append(str)
	}
	s.AsNums()
	s.Sort(1)
	act := s.strs
	if !slices.Equal(exp, act) {
		t.Fatalf("\nExpected:\n %v\nActual:\n %v", exp, act)
	}
}

func TestSortMonth(t *testing.T) {
	data := []string{"feb", "jul", "sep", "Jan", "JUN"}
	exp := []string{"Jan", "feb", "JUN", "jul", "sep"}

	s := New()
	for _, str := range data {
		s.Append(str)
	}
	s.AsMonths()
	s.Sort(1)
	act := s.strs
	if !slices.Equal(exp, act) {
		t.Fatalf("\nExpected:\n %v\nActual:\n %v", exp, act)
	}
}

func TestSortHumanReadableNumbers(t *testing.T) {
	data := []string{"100K", "20T", "32b", "3B", "2G", "10K"}
	exp := []string{"3B", "32b", "10K", "100K", "2G", "20T"}

	s := New()
	for _, str := range data {
		s.Append(str)
	}
	s.AsHumanNums()
	s.Sort(1)
	act := s.strs
	if !slices.Equal(exp, act) {
		t.Fatalf("\nExpected:\n %v\nActual:\n %v", exp, act)
	}
}

func TestSortByColumn(t *testing.T) {
	data := []string{"animal1:\tdog", "animal2:\tcat", "animal3:\tcow", "animal4:\trat"}
	exp := []string{"animal2:\tcat", "animal3:\tcow", "animal1:\tdog", "animal4:\trat"}

	s := New()
	for _, str := range data {
		s.Append(str)
	}
	s.Sort(2)
	act := s.strs
	if !slices.Equal(exp, act) {
		t.Fatalf("\nExpected:\n %v\nActual:\n %v", exp, act)
	}
}

func TestSortWithoutDuplicates(t *testing.T) {
	data := []string{"dog", "cat", "dog", "cow", "rat"}
	exp := []string{"cat", "cow", "dog", "rat"}

	s := New()
	for _, str := range data {
		s.Append(str)
	}
	s.Unique()
	s.Sort(1)
	act := s.strs
	if !slices.Equal(exp, act) {
		t.Fatalf("\nExpected:\n %v\nActual:\n %v", exp, act)
	}
}

func TestSortWithTrim(t *testing.T) {
	data := []string{"dog", "cat", "rat", "dog", "cow", " rat"}
	exp := []string{"cat", "cow", "dog", "dog", "rat", " rat"}

	s := New()
	for _, str := range data {
		s.Append(str)
	}
	s.IgnoreLeadingBlanks()
	s.Sort(1)
	act := s.strs
	if !slices.Equal(exp, act) {
		t.Fatalf("\nExpected:\n %v\nActual:\n %v", exp, act)
	}
}

func TestCheckIsSortedSuccess(t *testing.T) {
	data := []string{" rat", "cat", "cow", "dog", "dog", "rat"}
	exp := true

	s := New()
	for _, str := range data {
		s.Append(str)
	}
	s.IsSorted()
	s.Sort(1)
	act := s.isSorted
	if exp != act {
		t.Fatalf("\nExpected:\n %v\nActual:\n %v", exp, act)
	}
}

func TestCheckIsSortedFailed(t *testing.T) {
	data := []string{"dog", "cat", " rat", "dog", "cow", "rat"}
	exp := false

	s := New()
	for _, str := range data {
		s.Append(str)
	}
	s.IsSorted()
	s.Sort(1)
	act := s.isSorted
	if exp != act {
		t.Fatalf("\nExpected:\n %v\nActual:\n %v", exp, act)
	}
}
