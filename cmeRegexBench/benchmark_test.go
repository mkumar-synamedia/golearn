package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"testing"
)

func BenchmarkProcess(b *testing.B) {

	file, err := os.Open("samples.metrics")
	if err != nil {
		fmt.Printf("Error reading file: %v", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	defer file.Close()

	helpReg := regexp.MustCompile("^(#\\s+HELP\\s+([^\\s]+).*)$")
	typeReg := regexp.MustCompile("^(#\\s+TYPE\\s+([^\\s]+).*)$")
	sampleReg := regexp.MustCompile("^([^\\#][^\\s|{]*)\\s*({(.*)})?\\s+([^\\s]+)$")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, line := range lines {
			if helpReg.MatchString(line) {

			} else if typeReg.MatchString(line) {

			} else if sampleReg.MatchString(line) {

			}
		}
	}
}
