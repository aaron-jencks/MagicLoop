package magicloop

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/MagicTheGathering/mtg-sdk-go"
)

var cards []*mtg.Card

func GetValidCards(excludes []mtg.Card) []mtg.Card {
	result := make([]mtg.Card, 0, len(cards))
	for _, c := range cards {
		for _, e := range excludes {
			if e.MultiverseId != c.MultiverseId {
				result = append(result, *c)
			}
		}
	}
	return result
}

func loadCache(cacheLoc string) error {
	data, err := ioutil.ReadFile(cacheLoc)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &cards)
}

func saveCache(cacheLog string) error {
	fp, err := os.Create(cacheLog)
	if err != nil {
		return err
	}

	cenc, err := json.Marshal(cards)
	if err != nil {
		rerr := fp.Close()
		if rerr != nil {
			log.Default().Printf("Failed to close broken json cache: %s\n", rerr.Error())
		} else {
			rerr = os.Remove(cacheLog)
			if rerr != nil {
				log.Default().Printf("Failed to remove broken json cache: %s\n", rerr.Error())
			}
		}

		return err
	}

	_, err = fp.Write(cenc)
	if err != nil {
		rerr := fp.Close()
		if rerr != nil {
			log.Default().Printf("Failed to close broken json cache: %s\n", rerr.Error())
		} else {
			rerr = os.Remove(cacheLog)
			if rerr != nil {
				log.Default().Printf("Failed to remove broken json cache: %s\n", rerr.Error())
			}
		}

		return err
	}

	err = fp.Close()
	if err != nil {
		return err
	}

	return nil
}

func FetchCards(cacheLoc string) error {
	if FileExists(cacheLoc) {
		return loadCache(cacheLoc)
	} else {
		var err error
		cards, err = mtg.NewQuery().Where(mtg.CardSet, "MOM").All()
		if err != nil {
			return err
		}
		return saveCache(cacheLoc)
	}
}
