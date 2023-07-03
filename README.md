# MagicLoop
A software program to find infinite combos in magic the gathering.

## Note
This software just finds loops, it by no means indicates that you can pull them off, it's simply here to indicate combo synergy between cards.

#  How it works
The program loops over all of the cards in existence, for each one it simulates a custom game with up to `n` cards (you can specify `n` on the command line).

##  The custom game
To simulate the  game, some assumptions are made:
1. Assume ideal conditions for the player (deck is full, graveyard is empty, playing field is empty, mana cost is always satisfied for the first cast as well as first activation of abilities, all of the combo cards are in the player's hand)
2. The player can play any number of lands during their turn, there is no need to take extra turns during the simulation, unless it is for the purpose of taking infinite turns...
3. The Opponent always satisfies whatever the player needs, (always has cards in their hand, always has cards in the deck, always has permanents, they never have any responses, etc...)
4. The combo takes place on a single turn
5. The combo must be self-sustaining it must generate it's own mana, if needed, or untap mana, etc... The combo cannot abuse the assumptions to function.
6. The combo must reference the root card at least once per cycle (otherwise, why is the root card there?)

This allows for finding loops, even in ideal conditions.

## Simulating the stack
The game simulates the game, branching at every decision, using BFS to find loops.

## Storing results
Once a loop is found, it's inserted into a sql database, the data tables are as follows:
* cards:
  * id int
  * name string
* loops:
  * id int
  * root_card foreign key references cards:id
* loop_steps:
  * id int
  * loop_id foreign key references loops:id
  * step int
  * card_id foreign key references cards:id
  * action string
* loop_cards:
  * id int
  * loop_id foreign key references loops:id
  * card_id foreign key references cards:id
