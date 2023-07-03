package main

import (
	"flag"
	"log"

	"github.com/aaron-jencks/magicloop"
)

var comboSize int
var cardCache string
var forceUpdate bool
var fetchFileName string

func initializeFlags() {
	flag.IntVar(&comboSize, "n", 4, "set's the number of cards allowed in the combo, defaults to 4")
	flag.StringVar(&cardCache, "cache", "./cards.json", "the location of the downloaded set of cards should be stored between runs, defaults to ./cards.json")
	flag.BoolVar(&forceUpdate, "fetch", false, "indicates whether to force refresh the card cache from the server or not, defaults to false")
	flag.StringVar(&fetchFileName, "fetchfile", "SERVER", "tells where to read the scryfall data from, use SERVER to fetch from the internet, defaults to SERVER")
	flag.Parse()
}

func main() {
	initializeFlags()
	var err error
	if fetchFileName == "SERVER" {
		err = magicloop.FetchCards(cardCache, forceUpdate)
	} else {
		err = magicloop.FetchCardsFromFile(fetchFileName, cardCache, forceUpdate)
	}
	if err != nil {
		log.Fatalf("Failed to generate card cache: %s\n", err.Error())
		return
	}
}
