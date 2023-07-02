package magicloop

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/MagicTheGathering/mtg-sdk-go"
)

type Card struct {
	Id    string   // unique Id of the card
	Name  string   // the card name
	Types []string // the types of the card
	Text  string   // the oracle text of the card
}

func CreateCardFromSDK(c *mtg.Card) Card {
	return Card{
		Id:    string(c.Id),
		Name:  c.Name,
		Types: c.Types,
		Text:  c.Text,
	}
}

var cards []Card

func GetValidCards(excludes []Card) []Card {
	result := make([]Card, 0, len(cards))
	for _, c := range cards {
		for _, e := range excludes {
			if e.Id != c.Id {
				result = append(result, c)
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

func FetchCards(cacheLoc string, force bool) error {
	if !force && FileExists(cacheLoc) {
		log.Default().Println("Reading cache from disk.")
		return loadCache(cacheLoc)
	} else {
		log.Default().Println("Generating new cache on disk.")
		scards, err := mtg.NewQuery().Where(mtg.CardSet, "MOM").All()
		if err != nil {
			return err
		}
		for _, sc := range scards {
			cards = append(cards, CreateCardFromSDK(sc))
		}
		return saveCache(cacheLoc)
	}
}
