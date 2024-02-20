package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type job struct {
	id      int
	randmno int
}

type result struct {
	job        job
	sumofdigit int
}

var jobs = make(chan job, 10)
var results = make(chan result, 10)

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	return sum
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := result{job, digits(job.randmno)}
		results <- output
	}
	wg.Done()
}

func CreateWorker(noOfWorker int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorker; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func Allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomNO := rand.Intn(989)
		j := job{i, randomNO}
		jobs <- j

	}
	close(jobs)
}

func Results(done chan bool) {
	for result := range results {
		fmt.Printf("Job Id is %d, input Random no is %d, sum of digits: %d\n ", result.job.id, result.job.randmno, result.sumofdigit)
	}
	done <- true
}

func main() {
	startTime := time.Now()
	noOfJobs := 100
	go Allocate(noOfJobs)
	done := make(chan bool)
	go Results(done)
	noOfWorker := 20
	CreateWorker(noOfWorker)
	<-done
	endTime := time.Now().Sub(startTime).Seconds()
	fmt.Printf("Total Time Taken: %v seconds\n", endTime)

}
