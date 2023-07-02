# MagicLoop
A software program to find infinite combos in magic the gathering.

#  How it works
The program loops over all of the cards in existence, for each one it simulates a custom game with up to `n` cards (you can specify `n` on the command line).

##  The custom game
To simulate the  game, some assumptions are made:
1. The player's  deck is never empty
2. The player's graveyard starts empty
3. The Opponent always satisfies whatever the player needs, (always has cards in their hand, always has cards in the deck, always has permanents, etc...)

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
