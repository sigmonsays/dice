package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

type Options struct {
	// number of rolls
	Rolls int

	// number of dice to roll
	NumDice int
}

type Die int

func (d *Die) Roll() {
	v := (rand.Int() % 6) + 1
	*d = Die(v)
}
func NewRoll() *Die {
	d := new(Die)
	d.Roll()
	return d
}

type Dice []*Die

func (dice Dice) Roll() {
	for _, d := range dice {
		d.Roll()
	}
}
func (dice Dice) String() string {
	v := make([]int, len(dice))
	for i, d := range dice {
		v[i] = int(*d)
	}
	return fmt.Sprintf("%v", v)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	opts := &Options{
		Rolls:   1,
		NumDice: 1,
	}

	flag.IntVar(&opts.NumDice, "num", opts.NumDice, "number of dice to roll")
	flag.IntVar(&opts.NumDice, "n", opts.NumDice, "alias for num")
	flag.IntVar(&opts.Rolls, "rolls", opts.Rolls, "number of rolls")
	flag.IntVar(&opts.Rolls, "r", opts.Rolls, "alias for rolls")
	flag.Parse()

	for roll := 1; roll <= opts.Rolls; roll++ {
		dice := make(Dice, opts.NumDice)
		for i := 0; i < opts.NumDice; i++ {
			dice[i] = NewRoll()
		}
		fmt.Printf("roll %d %s\n", roll, dice)

	}
}
