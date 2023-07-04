package magicloop

type ICard interface {
	Play() IAction
	Sacrifice() IAction
	Discard() IAction
	ActivatedAbilities() []IAction
	HasTrigger(int) bool
	Trigger(int) IAction
	Morph() IAction
	Manifest() IAction
}

type IGameState interface {
	IClonable
	Step() []IGameState      // performs the current action and returns a list of next possible actions
	ValidActions() []IAction // returns the valid list of actions for the current state
}

type IAction interface {
	IClonable
	Act(IGameState) // performs this action on the current gamestate
	String() string // returns the string representation of this action
	Card() Card     // returns the card associated with this action
}

type IGame interface {
	IClonable
	Equals(IGame) bool // determines if one game is equal to another
}

/*
 * General procedure:
 * 1. Create some root state with all combo pieces in hand
 * 2. Call ValidActions() on the state to determine the set of possible moves
 * 3. Place these into a queue
 * 4. Pop an element off of the queue and perform the action
 * 5. If the resulting game state has already been visited, then we're in a loop, go to step 7.
 * 6. Go to step 2.
 * 7. Log the loop details
 * 8. Go to step 4.
 */
