package main

import (
	"flag"
	"log"
	"math"
)

/*
* Given an amount and a set of coins
* implement a function that outputs the
* optimal solution for the given amount
 */

// coin contains the name and value of a coin
type coin struct{
	name string
	value float64
}

var coins = []coin{
	{name: "1 pound", value: 1},
	{name: "50 pence", value: 0.50},
	{name: "20 pence", value: 0.20},
	{name: "10 pence", value: 0.10},
	{name: "5 pence", value: 0.05},
	{name: "1 penny", value:0.01},
}

//alculateChange return the coins required to calculare the change
func calculateChange(amount float64) map[coin]int{
	change := make(map[coin]int)
	for _, coin := range coins{
		if amount >= coin.value{
			count := math.Floor(amount / coin.value)
			amount = amount - count * coin.value
			change[coin] = int(count)
		}
	}
	return change
}

func printCoins(change map[coin]int){
	if len(change)==0{
		log.Println("No change found.")
		return
	}
	log.Println("Change has been calculated.")
	for coin, count := range change {
		log.Printf("%d x %s \n", count, coin.name)
	}
}

func main(){
	amount := flag.Float64("amount", 0.0, "The amount you want to make change")
	flag.Parse()
	change := calculateChange(*amount)
	printCoins(change)
}