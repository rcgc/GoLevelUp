package main

import (
	"flag"
	"log"
	"time"
)

/*
* Given a target date implement a function that
* outputs the numbers of nights until the
* given data
 */

var expectedFormat = "2006-01-02"

// parseTime validates and parses a given date string
func parseTime(target string) time.Time{
	pt, err := time.Parse(expectedFormat, target)
	if err != nil || time.Now().After(pt){
		log.Fatal("invalid date")
	}
	return pt
}

//calcSleeps returns the number of sleeps until the target
func calcSleeps(target time.Time) float64{
	return time.Until(target).Hours() / 24
}

func main(){
	bday := flag.String("bday", "", "Your next bday in YYYY-MM-DD format")
	flag.Parse()
	target := parseTime(*bday)
	log.Printf("You have %d sleeps until your birthday, Hurray!", int(calcSleeps(target)))
}