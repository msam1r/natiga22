package main

import (
	"os"

	"github.com/msam1r/natiga22/scrapper"
)

/* var (
	north  = "1001:33556"
	center = "35001:79525"
	south  = "80001:115593"
) */

func main() {
	file, _ := os.OpenFile("data.csv", os.O_RDWR|os.O_APPEND, 0775)
	defer file.Close()

	s := &scrapper.Scrapper{
		From: 1001,
		To:   33556,
		File: file,
	}

	s.Start()
}
