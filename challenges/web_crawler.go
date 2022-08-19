package challenges

import (
	"fmt"
	"sync"
)

var group = sync.WaitGroup{}

func crawl(workerName int, dataChan chan int) {
	for {
		select {
		case url := <-dataChan:
			fmt.Printf("Worker %d: Processing %d ...\n", workerName, url)
			fmt.Printf("Worker %d: Done!\n", workerName)
			group.Done()
		}

	}
}

func CrawlData(urls int) {
	dataChan := make(chan int)

	for i := 0; i < 5; i++ {
		go crawl(i, dataChan)
	}

	for i := 0; i < urls; i++ {
		group.Add(1)
		dataChan <- i
	}

	group.Wait()
}
