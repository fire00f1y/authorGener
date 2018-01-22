package main

import (
	"flag"
	"os"
	"io/ioutil"
	"log"
	"strings"
	"github.com/fire00f1y/authorGener/goodreads"
	"github.com/fire00f1y/authorGener/model"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	file := flag.String("file", "", "Abs or relative path to author csv file, including file name")
	key := flag.String("key", "", "Goodreads API key")
	runners := flag.Int("r", 10, "Number of runners to use querying the api")

	flag.Parse()
	if *file == "" || *key == "" {
		flag.Usage()
		os.Exit(1)
	}

	data, err := ioutil.ReadFile(*file)
	if err != nil {
		log.Fatalf("Failed to read from specified file [%s]: %v\n", *file, err)
	}

	authors := strings.Split(string(data), ",")
	counter := &model.Counter{}
	counter.Init()
	names := make(chan string, len(authors))
	finished := make(chan model.ChanStruct, *runners)

	fmt.Println("Starting runners")
	for i := 0; i < *runners; i++ {
		go runner(i, counter, *key, names, finished)
	}

	fmt.Printf("Starting processing for %d authors\n", len(authors))
	for _, author := range authors {
		names <- strings.TrimSpace(author)
	}
	close(names)

	finishedCount := 0
	unknowns := make([]string, 0)
	for {
		s := <-finished

		fmt.Printf("Runner finished processing %d\nWas unable to process these:\n", s.Processed)
		for _, u := range s.Unknowns {
			unknowns = append(unknowns, u)
		}

		finishedCount += 1
		if finishedCount == *runners {
			break
		}
	}

	printUnknowns(unknowns)
	length := time.Since(start)
	fmt.Printf("Aggregation lasted %v\n", length)
	counter.Print()
}

func runner(runnerId int, counter *model.Counter, key string, names chan string, finished chan model.ChanStruct) {
	unknowns := make([]string, 0)
	processed := 0
	for author := range names {
		fmt.Printf("[%d] Processing %s\n", runnerId, author)
		id, err := goodreads.GetAuthorId(author, key, false)
		if id == "" {
			correctedName := goodreads.CorrectedName(author)
			fmt.Printf("No name or error returned for [%s]. Trying with corrected name [%s].\n", author, correctedName)
			id, err = goodreads.GetAuthorId(correctedName, key, true)
		}
		if err != nil {
			fmt.Printf("[%d] Error while getting id for author [%s]: %v\n", runnerId, author, err)
			counter.AddGender("unknown")
			continue
		}
		if id == "" {
			unknowns = append(unknowns, author)
			continue
		}
		gender, err := goodreads.GetAuthorInfo(id, key)
		if err != nil {
			fmt.Printf("[%d] Error while getting gender for author [%s]:[%s]: %v\n", runnerId, id, author, err)
			unknowns = append(unknowns, author)
		}
		if gender == "" {
			unknowns = append(unknowns, author)
		}
		processed += 1
		counter.AddGender(gender)
	}
	fmt.Printf("Finished runner %d\n", runnerId)

	finished <- model.ChanStruct{Processed:processed, Unknowns: unknowns}
}

func printUnknowns(unknowns []string) {
	if len(unknowns) == 0 {
		return
	}
	fmt.Printf("----------------- %d unknowns -----------------\n", len(unknowns))
	for _, n := range unknowns {
		fmt.Printf("\t%s\n", n)
	}
	fmt.Println("-----------------------------------------------")
}

