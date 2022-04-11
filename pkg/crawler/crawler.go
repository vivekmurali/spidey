package crawler

import (
	"fmt"
	"sync"
)

func GetPage(url string, wg *sync.WaitGroup, count int) {
	fmt.Println(url, count)

	// Create new waitgroup before sending all links here
	// Recursive function

	wg.Done()
}
