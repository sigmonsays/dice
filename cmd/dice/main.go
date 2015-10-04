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

	// sides of a die
	Sides int
}

type Die struct {
	Sides int
	Value int
}

func NewDie() *Die {
	return &Die{
		Sides: 6,
		Value: 0,
	}
}

func NewSidedDie(sides int) *Die {
	d := NewDie()
	d.Sides = sides
	return d
}

func (d *Die) Roll() {
	v := (rand.Int() % d.Sides) + 1
	d.Value = v
}

// make a new die and roll it
func NewRoll(sides int) *Die {
	d := NewSidedDie(sides)
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
		v[i] = d.Value
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
		Sides:   6,
	}

	flag.IntVar(&opts.NumDice, "num", opts.NumDice, "number of dice to roll")
	flag.IntVar(&opts.NumDice, "n", opts.NumDice, "alias for num")
	flag.IntVar(&opts.Rolls, "rolls", opts.Rolls, "number of rolls")
	flag.IntVar(&opts.Rolls, "r", opts.Rolls, "alias for rolls")
	flag.IntVar(&opts.Sides, "s", opts.Sides, "alias for sides")
	flag.IntVar(&opts.Sides, "sides", opts.Sides, "alias for sides")
	flag.Parse()

	for roll := 1; roll <= opts.Rolls; roll++ {
		dice := make(Dice, opts.NumDice)
		for i := 0; i < opts.NumDice; i++ {
			dice[i] = NewRoll(opts.Sides)
		}
		fmt.Printf("roll %d %s\n", roll, dice)

	}
}
