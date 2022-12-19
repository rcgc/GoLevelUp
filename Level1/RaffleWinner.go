package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"
)

/*
* Given an input.json file with raggle entries
* implement a function that outputs a list of
* structs containins the entries
 */

const path = "entries.json"

// raffleEntry is the struct we unmarshal raffle entries into
type raffleEntry struct{
	ID string `json:"id"`
	Name string `json:"name"`
}

// importData reads the raffle entries from file and creates the entries slice
func importData() []raffleEntry{
	file, err := os.ReadFile(path)
	if err != nil{
		log.Fatal(err)
	}
	var data[] raffleEntry
	err = json.Unmarshal(file, &data)
	if err != nil{
		log.Fatal(err)
	}
	return data
}

// getWinner returns a random winner from a slice of raffle entries
func getWinner(entries []raffleEntry) raffleEntry{
	rand.Seed(time.Now().Unix())
	wi := rand.Intn(len(entries))
	return entries[wi]
}

func main(){
	entries := importData()
	log.Println("And... the raffle winning entry is...")
	winner := getWinner(entries)
	time.Sleep(500 * time.Millisecond)
	log.Println(winner)
}