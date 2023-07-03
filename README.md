# dice

simple cli utility to roll the dice

    Usage of dice:
     -n=3: alias for num
     -num=3: number of dice to roll
     -r=3: alias for rolls
     -roll-until="": roll until the sequence is met (coma delimited)
     -rolls=3: number of rolls
     -s=6: alias for sides
     -sides=6: alias for sides

Roll some dice

      $ dice
      roll 1: 1 2 1 (4 face value)
      roll 2: 6 6 4 (16 face value)
      roll 3: 1 5 2 (8 face value)


Roll until a sequence

      $ dice --roll-until 1,2,3,4,5,6
      rolling until sequence [1 2 3 4 5 6]
      rolled [1 2 3 4 5 6] in 407 rolls


Roll until specific numbers come up

      $ dice --roll-until 1,1,1
