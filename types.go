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

func (g Game) Clone() interface{} {
	var result Game = Game{
		graveyard:     make([]Card, len(g.graveyard)),
		exile:         make([]Card, len(g.exile)),
		mapped_exiles: map[string][]Card{},
		counters:      map[int]int{},
		stack:         make([]IAction, len(g.stack)),
		mana:          map[int]int{},
	}

	for gi := range g.graveyard {
		result.graveyard[gi] = g.graveyard[gi].Clone().(Card)
	}
	for ei := range g.exile {
		result.exile[ei] = g.exile[ei].Clone().(Card)
	}
	for si := range g.stack {
		result.stack[si] = g.stack[si].Clone().(IAction)
	}

	for k, v := range g.mapped_exiles {
		result.mapped_exiles[k] = make([]Card, len(v))
		for vi := range v {
			result.mapped_exiles[k][vi] = v[vi].Clone().(Card)
		}
	}

	for k, v := range g.counters {
		result.counters[k] = v
	}
	for k, v := range g.mana {
		result.mana[k] = v
	}

	return result
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

func (gs GameState) Step() []IGameState {
	gs.current_action.Act(gs)
	actions := gs.ValidActions()

	var result []IGameState
	for _, act := range actions {
		gsc := gs.Clone().(GameState)
		gsc.current_action = act
		result = append(result, gsc)
	}
	return result
}

func (gs GameState) ValidActions() []IAction {
	return nil
}

func (gs GameState) Clone() interface{} {
	var result GameState = GameState{
		current_action:  gs.current_action,
		game:            gs.game,
		previous_states: make([]IGame, len(gs.previous_states)),
		actions:         make([]IAction, len(gs.actions)),
		cards:           make([]Card, len(gs.cards)),
	}
	for psi := range gs.previous_states {
		result.previous_states[psi] = gs.previous_states[psi].Clone().(IGame)
	}
	for ai := range gs.actions {
		result.actions[ai] = gs.actions[ai].Clone().(IAction)
	}
	for ci := range gs.cards {
		result.cards[ci] = gs.cards[ci].Clone().(Card)
	}
	return result
}

type Simulation struct {
	action_queue []IGameState
}
