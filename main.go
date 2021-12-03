package main

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
)



func main() {
	thing := "films"
	number := 1
	
	s := spinner.New(spinner.CharSets[2], 100*time.Millisecond)
	fmt.Printf(" Fetching info for: %s/%d\n", thing, number)
	// s.Suffix = suffix
	s.Color("blue")
	s.FinalMSG = "Done!\n"
	s.Start()

	flimInfo := GetFilm(number)
	s.Stop()

	fmt.Println(flimInfo.Director)
}