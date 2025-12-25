package cmd

import (
	"errors"
	"os"
	"os/exec"
	"testing"
)

func TestSortCmd_WithSingleFlag_Success(t *testing.T) {
	data := []byte("10\n1\n2")

	exp := "1: 1\n2: 2\n3: 10\n"
	act, err := execSort(t, data, "-n")
	if err != nil {
		t.Fatal(err.Error())
	}

	if exp != string(act) {
		t.Errorf("\nExpected:\n%q\n\nActual:\n%q", exp, act)
	}
}

func TestSortCmd_WithCombinedFlags_Success(t *testing.T) {
	data := []byte("10\n1\n2")

	exp := "3: 10\n2: 2\n1: 1\n"
	act, err := execSort(t, data, "-nr")
	if err != nil {
		t.Fatal(err.Error())
	}

	if exp != string(act) {
		t.Errorf("\nExpected:\n%q\n\nActual:\n%q", exp, act)
	}
}

func TestSortCmd_WithInvalidCombinedFlags_Error(t *testing.T) {
	data := []byte("10\n1\n2")

	exp := "sort: unknown sort type\nexit status 1\n"
	act, actErr := execSort(t, data, "-nH")

	if exp != string(act) {
		t.Errorf("\nExpected:\n%q\n\nActual:\n%q", exp, act)
	}

	if actErr == nil {
		t.Fatal("expected error, got nil")
	}

	var exitErr *exec.ExitError
	if !errors.As(actErr, &exitErr) {
		t.Fatalf("expected ExitError, got %T", actErr)
	}

	if exitErr.ExitCode() != 1 {
		t.Fatalf("expected exit code 1, got %d", exitErr.ExitCode())
	}
}

func execSort(t *testing.T, data []byte, flags string) ([]byte, error) {
	tmpfile, err := os.CreateTemp("", "test_input_*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(data); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	cmd := exec.Command("go", "run", "../main.go", flags, tmpfile.Name())
	output, err := cmd.CombinedOutput()

	return output, err
}
