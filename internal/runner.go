package internal

import (
	"time"
	"sync"
	"net/http"
)
func RunTest(cfg Config) ([]Result, time.Duration) {
	var wg sync.WaitGroup
	sem := make(chan struct{}, cfg.Concurrency)
	results := make([]Result, cfg.Requests)

	start := time.Now()

	for i := 0; i < cfg.Requests; i++ {
		wg.Add(1)
		go func (idx int)  {
			defer wg.Done()
			sem <- struct{}{}
			resp, err := http.Get(cfg.URL)
			if err != nil {
				results[idx] = Result{StatusCode: 0, Error: err}
			} else {
				results[idx] = Result{StatusCode: resp.StatusCode, Error: nil}
				resp.Body.Close()
			}
			<-sem
		}(i)
	}
	wg.Wait()
	duration := time.Since(start)
	return results, duration
}