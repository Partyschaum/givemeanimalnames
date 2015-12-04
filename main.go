package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"time"
)

const animalsFile = "animals.txt"
const adjectivesFile = "adjectives.txt"

var animals []string
var animal string
var adjectives []string

var n = flag.Bool("n", false, "omit trailing newline")
var alliterate = flag.Bool("a", false, "alliterate")
var sep = flag.String("s", " ", "separator")

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	flag.Parse()

	animals = fileToSortedStringSlice(animalsFile)
	adjectives = fileToSortedStringSlice(adjectivesFile)

	adjective, err := randomFromList(adjectives)
	if err != nil {
		log.Fatalf("givemeanimalnames: %v", err)
	}

	if *alliterate {
		group, err := subgroupFromList(adjective, animals)
		if err == nil {
			animal = group[random.Intn(len(group))]
		}
	}

	if animal == "" {
		animal = animals[random.Intn(len(animals))]
	}

	name := fmt.Sprintf("%s%s%s",
		adjective,
		*sep,
		animal,
	)

	if !*n {
		fmt.Println(name)
	} else {
		fmt.Print(name)
	}
}

func subgroupFromList(char string, list []string) ([]string, error) {
	var group []string

	for _, j := range list {
		if char[0] == j[0] {
			group = append(group, j)
		}
	}

	if len(group) > 0 {
		return group, nil
	}
	return group, errors.New("no words found for subgroup!")
}

func randomFromList(list []string) (string, error) {
	if len(list) == 0 {
		return "", errors.New("list of words is empty!")
	}

	return list[random.Intn(len(list))], nil
}

func fileToSortedStringSlice(file string) (out []string) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("givemeanimalnames: %v", err)
	}

	input := bufio.NewScanner(f)
	for input.Scan() {
		out = append(out, input.Text())
	}

	sort.Strings(out)

	f.Close()
	return
}
