package concurrency

type WebsiteChecker func(string) bool

type result struct {
	url    string
	result bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultsChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultsChannel <- result{u, wc(u)}
			// results[u] = wc(u)
		}(url)

	}

	for i := 0; i < len(urls); i++ {
		r := <-resultsChannel
		results[r.url] = r.result
	}

	return results
}
