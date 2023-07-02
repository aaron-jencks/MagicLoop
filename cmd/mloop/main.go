package main

import (
	"flag"
	"log"

	"github.com/aaron-jencks/magicloop"
)

var comboSize int
var cardCache string
var forceUpdate bool

func initializeFlags() {
	flag.IntVar(&comboSize, "n", 4, "set's the number of cards allowed in the combo, defaults to 4")
	flag.StringVar(&cardCache, "cache", "./cards.json", "the location of the downloaded set of cards should be stored between runs, defaults to ./cards.json")
	flag.BoolVar(&forceUpdate, "fetch", false, "indicates whether to force refresh the card cache from the server or not, defaults to false")
	flag.Parse()
}

func main() {
	initializeFlags()
	err := magicloop.FetchCards(cardCache, forceUpdate)
	if err != nil {
		log.Fatalf("Failed to generate card cache: %s\n", err.Error())
		return
	}
}
