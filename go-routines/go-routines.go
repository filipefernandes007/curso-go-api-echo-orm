package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

/*
func goRoutine() {
	time.Sleep(1 * time.Second)
	fmt.Println("In go routine")
}
 */

type AirlinesType struct {
	Airlines []string
}

func NewAirlinesType() *AirlinesType {
	return &AirlinesType{Airlines: make([]string, 0)}
}

func callAirlines(id int) (string, error) {
	resp, err := http.Get(
		fmt.Sprintf("https://api.instantwebtools.net/v1/airlines/%d", id),
	)
	if err != nil {
		return "", errors.New("could not call airlines")
	}
	defer resp.Body.Close()

	// Read body from response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("could not read response")
	}

	return string(body), nil
}

func gotRoutine(wg *sync.WaitGroup, id int, airlines *AirlinesType) {
	defer wg.Done()

	fmt.Printf("Starting for %d\n", id)
	response, err := callAirlines(id)
	if err != nil {
		log.Fatal("Error reading response", err)
	}

	a := airlines.Airlines
	a = append(a, response)
	airlines.Airlines = a
	fmt.Printf("Ended for %v\n", id)
}

func main() {
	var wg sync.WaitGroup

	//var airlines []string = make([]string, 0)
	var values = []int{1,2,3,4,5}
	airlines := NewAirlinesType()
	// go goRoutine()
	start := time.Now()
	for i, id := range values {
		fmt.Printf("Starting %d\n", i)
		wg.Add(1)
		go gotRoutine(&wg, id, airlines)
		//response, err := callAirlines(id)
		/*
		if err != nil {
			log.Fatal("Error calling airlines. ", err)
		}

		 */
		//airlines = append(airlines, response)
	}
	wg.Wait()
	duration := time.Since(start)

	fmt.Printf("%s\n", airlines)
	fmt.Println("Duration ", duration)
}
