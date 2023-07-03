package magicloop

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	scryfall "github.com/BlueMonday/go-scryfall"
)

type CardFace struct {
	Name     string // the name of the face
	ManaCost string // the mana cost of the face, if any
	Type     string // the type string of the face, if any
	Text     string // the printed text for this face, if any
}

func (cf CardFace) Equals(o CardFace) bool {
	return cf.Name == o.Name &&
		cf.ManaCost == o.ManaCost &&
		cf.Type == o.Type &&
		cf.Text == o.Text
}

type Card struct {
	Id            string           // unique Id of the card
	Name          string           // the card name
	Type          string           // the types of the card
	Text          string           // the oracle text of the card
	Layout        scryfall.Layout  // for detecting flip cards
	ManaCost      string           // the mana cost of the card
	ColorIdentity []scryfall.Color // the color identity of the card
	CardFaces     []CardFace       // the faces of this card
	CardFace      int              // the current face of the card
	Counters      map[int]int      // a map of the counters placed on this card
}

func (c Card) Equals(o Card) bool {
	if len(c.ColorIdentity) != len(o.ColorIdentity) ||
		len(c.CardFaces) != len(o.CardFaces) {
		return false
	}

	for _, ci := range c.ColorIdentity {
		found := false
		for _, oci := range o.ColorIdentity {
			if ci == oci {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	for _, ci := range c.CardFaces {
		found := false
		for _, oci := range o.CardFaces {
			if ci.Equals(oci) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return c.Id == o.Id &&
		c.Name == o.Name &&
		c.Type == o.Type &&
		c.Text == o.Text &&
		c.Layout == o.Layout &&
		c.ManaCost == o.ManaCost
}

func CreateCardFromSDK(c scryfall.Card) Card {
	var cardfaces []CardFace = nil
	if len(c.CardFaces) > 0 {
		cardfaces = make([]CardFace, 0, len(c.CardFaces))
		for _, cf := range c.CardFaces {
			cardfaces = append(cardfaces, CardFace{
				Name:     cf.Name,
				ManaCost: cf.ManaCost,
				Type:     cf.TypeLine,
				Text:     *cf.OracleText,
			})
		}
	}
	return Card{
		Id:            c.OracleID,
		Name:          c.Name,
		Type:          c.TypeLine,
		Text:          c.OracleText,
		Layout:        c.Layout,
		ManaCost:      c.ManaCost,
		ColorIdentity: c.ColorIdentity,
		CardFaces:     cardfaces,
		Counters:      map[int]int{},
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
		err := loadCache(cacheLoc)
		if err != nil {
			return err
		}
		log.Default().Printf("Found %d cards\n", len(cards))
		return nil
	} else {
		log.Default().Println("Generating new cache on disk.")

		client, err := scryfall.NewClient()
		if err != nil {
			return err
		}

		sco := scryfall.SearchCardsOptions{
			Unique:        scryfall.UniqueModeCards,
			Order:         scryfall.OrderColor,
			Dir:           scryfall.DirDesc,
			IncludeExtras: false,
		}

		scards, err := client.SearchCards(context.Background(), "set:mom", sco)
		if err != nil {
			return err
		}

		for _, sc := range scards.Cards {
			cards = append(cards, CreateCardFromSDK(sc))
		}

		log.Default().Printf("Found %d cards\n", len(cards))
		return saveCache(cacheLoc)
	}
}

func FetchCardsFromFile(filename, cacheLoc string, force bool) error {
	if !force && FileExists(cacheLoc) {
		log.Default().Println("Reading cache from disk.")
		err := loadCache(cacheLoc)
		if err != nil {
			return err
		}
		log.Default().Printf("Found %d cards\n", len(cards))
		return nil
	} else {
		log.Default().Println("Generating new cache on disk.")

		data, err := ioutil.ReadFile(filename)
		if err != nil {
			return err
		}

		var scards []scryfall.Card
		err = json.Unmarshal(data, &scards)
		if err != nil {
			return err
		}

		for _, sc := range scards {
			cards = append(cards, CreateCardFromSDK(sc))
		}

		log.Default().Printf("Found %d cards\n", len(cards))
		return saveCache(cacheLoc)
	}
}
