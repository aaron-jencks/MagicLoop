package magicloop

import (
	"flag"
	"log"
)

var comboSize int
var cardCache string

func initializeFlags() {
	flag.IntVar(&comboSize, "n", 4, "set's the number of cards allowed in the combo, defaults to 4")
	flag.StringVar(&cardCache, "cache", "./cards.json", "the location of the downloaded set of cards should be stored between runs, defaults to ./cards.json")
	flag.Parse()
}

func main() {
	initializeFlags()
	err := FetchCards(cardCache)
	if err != nil {
		log.Fatalf("Failed to generate card cache: %s\n", err.Error())
		return
	}
}
