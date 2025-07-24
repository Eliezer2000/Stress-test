package main

import (
	"github.com/Eliezer2000/Stress-test/internal"
)

func main() {
	cfg := internal.ParseFlags()
	results, duration := internal.RunTest(cfg)
	internal.PrintReport(results, duration)
}