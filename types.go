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
	deck_count    int
	life          int
	graveyard     []Card
	exile         []Card
	mapped_exiles map[string][]Card
	counters      map[int]int
	stack         []IAction
	mana          map[int]int
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
