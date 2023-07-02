package magicloop

import "github.com/MagicTheGathering/mtg-sdk-go"

type IGameState interface {
	Step() []IGameState      // performs the current action and returns a list of next possible actions
	ValidActions() []IAction // returns the valid list of actions for the current state
}

type IAction interface {
	Act(IGameState) // performs this action on the current gamestate
	String() string // returns the string representation of this action
	Card() mtg.Card // returns the card associated with this action
}

type IGame interface {
	Equals(IGame) bool // determines if one game is equal to another
}
