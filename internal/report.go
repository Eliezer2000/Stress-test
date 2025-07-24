package internal

import (
	"time"
	"fmt"
)

func PrintReport(results []Result, duration time.Duration) {
	total := len(results)
	success := 0
	statusCount := make(map[int]int)
	errorCount := 0

	for _, r := range results {
		if r.Error != nil {
			errorCount++
			statusCount[0]++
		} else {
			if r.StatusCode == 200 {
				success++
			}
			statusCount[r.StatusCode]++
		}
	}
	fmt.Println("====== Stress Test Report ======")
    fmt.Printf("Total time: %v\n", duration)
    fmt.Printf("Total requests: %d\n", total)
    fmt.Printf("HTTP 200: %d\n", success)
    fmt.Printf("Errors: %d\n", errorCount)
    fmt.Println("Status code distribution:")

	for code, count := range statusCount {
		if code == 0 {
			fmt.Printf(" Errors (no response): %d\n", count)
		} else {
			fmt.Printf(" %d: %d\n", code, count)
		}
	}
	fmt.Println("================================")
}