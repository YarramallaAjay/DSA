package workerpool

import (
	"fmt"
	"sync"
)

type Job struct {
	Id    int
	Value int
}

type Result struct {
	Id     int
	Square int
}

func Worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {

	defer wg.Done()

	for j := range jobs {
		fmt.Printf("worker: %v is working with job : %v \n", id, j.Id)
		val := j.Value
		results <- Result{
			Id:     j.Id,
			Square: val * 2,
		}

	}
}

func CollectResults(Results <-chan Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for r := range Results {
		fmt.Printf("Job %v completed and received the response as: %v \n", r.Id, r.Square)
	}
}

func Dispatcher(jobsCount int, workersCount int) {

	jobs := make(chan Job)
	results := make(chan Result)

	var wg sync.WaitGroup

	wg.Add(workersCount)

	for i := 0; i < workersCount; i++ {

		go Worker(i, jobs, results, &wg)

	}

	var resultWg sync.WaitGroup

	resultWg.Add(1)

	go CollectResults(results, &resultWg)

	for i := 0; i <= jobsCount; i++ {
		jobs <- Job{
			Id:    i,
			Value: i + 1,
		}
	}

	close(jobs)

	wg.Wait()

	close(results)

	resultWg.Wait()

}
