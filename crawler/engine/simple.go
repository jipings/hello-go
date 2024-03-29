package engine

import (
	"log"
)

type SimpleEngine struct {}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		if isDuplicate(r.Url) {
			continue
		}
		requests = append(requests, r)
	}

	for len(requests) >0 {
		r := requests[0]
		requests = requests[1:]
		if isDuplicate(r.Url) {
			continue
		}
		parseResult, err := Worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Item %v", item)
		}
	}
}