package io

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestReadFile(t *testing.T) {
	path, err := filepath.Abs("../database/common.go")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(path)

	file, err := os.Open("test.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		t.Fatal(err)
	}

	io.Copy(os.Stdout, bytes.NewReader(content))
	os.Stdout.WriteString("\n")
}
