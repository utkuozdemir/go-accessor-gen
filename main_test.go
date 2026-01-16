package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	// Setup
	inputContent := `package testdata

import "time"

type Simple struct {
	Name *string
	Age  *int
}

type Complex struct {
	Tags    *[]string
	Timeout *time.Duration
	// Should not generate for these
	Normal string
	Slice  []int
}
`
	dir := t.TempDir()
	inputFile := filepath.Join(dir, "input.go")
	if err := os.WriteFile(inputFile, []byte(inputContent), 0644); err != nil {
		t.Fatalf("failed to write input file: %v", err)
	}

	// Run
	args := []string{inputFile}
	suffix := ".gen.go"
	if err := Run(args, suffix); err != nil {
		t.Fatalf("Run failed: %v", err)
	}

	// Verify
	outputFile := filepath.Join(dir, "input.gen.go")
	content, err := os.ReadFile(outputFile)
	if err != nil {
		t.Fatalf("failed to read output file: %v", err)
	}
	output := string(content)

	expectedChecks := []string{
		"func (s *Simple) GetName() string",
		"func (s *Simple) SetName(v string)",
		"func (s *Complex) GetTags() []string",
		"func (s *Complex) SetTags(v []string)",
		"func (s *Complex) GetTimeout() time.Duration",
		"func (s *Complex) SetTimeout(v time.Duration)",
	}

	for _, check := range expectedChecks {
		if !strings.Contains(output, check) {
			t.Errorf("output missing expected code: %s", check)
		}
	}

	unexpectedChecks := []string{
		"GetNormal",
		"SetNormal",
		"GetSlice",
		"SetSlice",
	}

	for _, check := range unexpectedChecks {
		if strings.Contains(output, check) {
			t.Errorf("output contains unexpected code: %s", check)
		}
	}
}
