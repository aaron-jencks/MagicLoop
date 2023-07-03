package magicloop

const (
	MANA_WILD = iota
	MANA_COLORLESS
	MANA_GREEN
	MANA_RED
	MANA_BLUE
	MANA_WHITE
	MANA_BLACK
)

const (
	COUNTERS_POISION = iota
	COUNTERS_DIVINITY
	COUNTERS_LOYALTY
	COUNTERS_ENERGY
	COUNTERS_OIL
	COUNTERS_P1P1
	COUNTERS_N1N1
	COUNTERS_P1N1
	COUNTERS_N1P1
)

type Game struct {
	graveyard     []Card
	exile         []Card
	mapped_exiles map[string][]Card
	counters      map[int]int
	stack         []IAction
	mana          map[int]int
}

func (g Game) Equals(other IGame) bool {
	og, ok := other.(Game)
	if !ok {
		return false
	}

	if len(g.graveyard) != len(og.graveyard) ||
		len(g.exile) != len(og.exile) ||
		len(g.mapped_exiles) != len(og.mapped_exiles) ||
		len(g.counters) != len(og.counters) ||
		len(g.mana) != len(og.mana) {
		return false
	}

	for _, gc := range g.graveyard {
		found := false
		for _, ogc := range og.graveyard {
			if gc.Equals(ogc) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	for _, ec := range g.exile {
		found := false
		for _, oec := range og.exile {
			if ec.Equals(oec) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	for k, v := range g.mapped_exiles {
		ov, ok := og.mapped_exiles[k]
		if !ok {
			return false
		}

		for _, ec := range v {
			found := false
			for _, oec := range ov {
				if ec.Equals(oec) {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}

	for k, v := range g.counters {
		ov, ok := og.counters[k]
		if !ok || v != ov {
			return false
		}
	}

	for k, v := range g.mana {
		ov, ok := og.mana[k]
		if !ok || v != ov {
			return false
		}
	}

	return true
}

type GameState struct {
	current_action  IAction
	game            IGame
	previous_states []IGame
	actions         []IAction
	cards           []Card
}

type Simulation struct {
	action_queue []IGameState
}
