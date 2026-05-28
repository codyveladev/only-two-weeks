package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Result struct {
	URL        string
	StatusCode int
	SizeBytes  int
	Err        error
}

func CreateUrls() []string {
	urls := []string{}
	for i := 0; i < 10; i++ {
		index := strconv.Itoa(i)
		url := "https://jsonplaceholder.typicode.com/todos/" + index
		urls = append(urls, url)
	}
	return urls
}

func TimeGetsWithoutGoRoutines() {
	urls := CreateUrls()
	result := []Result{}
	start := time.Now()
	for _, url := range urls {
		response, err := http.Get(url)
		if err != nil {
			result = append(result, Result{URL: url, Err: err})
			continue
		}
		body, err := io.ReadAll(response.Body)
		if err != nil {
			result = append(result, Result{URL: url, StatusCode: response.StatusCode, Err: err})
			continue
		}
		result = append(result, Result{
			URL:        url,
			StatusCode: response.StatusCode,
			SizeBytes:  len(body),
			Err:        nil,
		})

	}

	elapsed := time.Since(start)
	fmt.Println("gets without go routines took: ", elapsed)
	for _, r := range result {
		fmt.Println(r)
	}
}

func main() {
	urls := CreateUrls()
	results := make(chan Result, len(urls))

	var wg sync.WaitGroup
	start := time.Now()
	for _, url := range urls {
		wg.Add(1)

		go func(u string) {
			defer wg.Done()
			response, err := http.Get(u)
			if err != nil {
				results <- Result{URL: u, Err: err}
				return
			}
			defer response.Body.Close()
			body, err := io.ReadAll(response.Body)
			if err != nil {
				results <- Result{URL: u, StatusCode: response.StatusCode, Err: err}
				return
			}
			result := Result{
				URL:        u,
				StatusCode: response.StatusCode,
				SizeBytes:  len(body),
				Err:        nil,
			}
			results <- result

		}(url)
	}

	wg.Wait()

	close(results)
	elapsed := time.Since(start)
	fmt.Println("Go routines executed in :", elapsed)
	for r := range results {
		fmt.Println(r)
	}
	TimeGetsWithoutGoRoutines()
}
