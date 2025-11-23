package workerpool2

import (
	"fmt"
	"sync"
)

type Image struct {
	Id   int
	Url  string
	Size []int
}

type Output struct {
	Image Image
}

type ImageJob struct {
	Id     int
	Action string
	Image  Image
}

func Enlarge(image Image) Image {
	image.Size[0] += 100
	image.Size[1] += 100
	return image
}

func Shrink(image Image) Image {
	image.Size[0] -= 100
	image.Size[1] -= 100
	return image
}

func ImageWorker(id int, jobs <-chan ImageJob, out chan<- Output, wg *sync.WaitGroup) {

	defer wg.Done()
	for job := range jobs {
		fmt.Printf("worker: %v is working with job : %v \n", id, job.Id)
		var processedImage Image
		switch job.Action {
		case "enlarge":
			processedImage = Enlarge(job.Image)
		case "shrink":
			processedImage = Shrink(job.Image)

		}
		out <- Output{Image: processedImage}
	}
}

func Collector(out <-chan Output, wg *sync.WaitGroup) {

	defer wg.Done()

	for o := range out {
		fmt.Println("Processed Image ID:", o.Image.Id,
			" New Size:", o.Image.Size)
	}
}

func ImageDispatcher(jobsCount, WorkersCount int) {

	Jobs := make(chan ImageJob)
	Output := make(chan Output)

	var wg sync.WaitGroup

	wg.Add(WorkersCount)

	for i := 0; i < WorkersCount; i++ {
		go ImageWorker(i, Jobs, Output, &wg)
	}

	var outputWg sync.WaitGroup

	go Collector(Output, &outputWg)

	for i := 0; i < jobsCount; i++ {

		image := Image{
			Id:   i,
			Url:  fmt.Sprintf("new Image: %v", i),
			Size: []int{i * 5, i * 10},
		}
		job := ImageJob{
			Id:     i,
			Action: "enlarge",
			Image:  image,
		}

		Jobs <- job
	}

	close((Jobs))
	wg.Wait()

	close(Output)
	outputWg.Wait()
}
