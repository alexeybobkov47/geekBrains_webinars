package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

func main() {
	str := "html"
	sites := []string{
		"https://yandex.ru",
		"https://golang.org",
		"https://google.com",
		"https://github.com",
	}

	var (
		result = make([]string, 0, len(sites))
		errs   int
	)

	// you can add flag here
	if !true {
		result, errs = search(str, sites)
	} else {
		result, errs = searchConcurrency(str, sites)
	}

	if errs > 0 {
		log.Printf("There are %v errors during request", errs)
	}

	if len(result) == 0 {
		log.Println("empty result")
		return
	}

	log.Printf("Sites: %v", result)
}

// the first version of search()
func search(str string, sites []string) ([]string, int) {
	out := make([]string, 0, 1)
	errs := 0

	for _, site := range sites {
		res, err := getReq(site)
		if err != nil {
			errs++
			log.Print(err)
			continue
		}

		if strings.Contains(string(res), str) {
			out = append(out, site)
		}
	}

	return out, errs
}

func getReq(reqURL string) ([]byte, error) {
	resp, err := http.Get(reqURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type chunk struct {
	site string
	err  error
}

// the second version of search
func searchConcurrency(str string, sites []string) ([]string, int) {
	wg := sync.WaitGroup{}
	results := make(chan chunk)

	for _, site := range sites {
		wg.Add(1)
		go func(site string) {
			defer wg.Done()
			res, err := getReq(site)
			if err != nil {
				results <- chunk{site: site, err: err}
				return
			}

			if strings.Contains(string(res), str) {
				results <- chunk{site: site}
			}
		}(site)
	}

	go func() {
		wg.Wait()
		log.Println("channel closed")
		close(results)
	}()

	out := make([]string, 0, len(sites))
	errs := 0
	for chunk := range results {
		if chunk.err != nil {
			errs++
			// log.Printf("Error %v for site %v", chunk.err, chunk.site)
			continue
		}

		// log.Println(chunk.site)
		out = append(out, chunk.site)
	}

	return out, errs
}
