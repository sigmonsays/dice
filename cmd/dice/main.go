package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
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

type Dicer interface {
	Roll()
	FaceValue() int
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

func (d *Die) FaceValue() int {
	return d.Value
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
	v := make([]string, len(dice))

	char_width := len(fmt.Sprintf("%d", dice[0].Sides))
	f := "%" + fmt.Sprintf("%d", char_width) + "d"

	for i, d := range dice {
		v[i] = fmt.Sprintf(f, d.Value)
	}
	return strings.Join(v, " ")
}

func (dice Dice) FaceValue() (value int) {
	for _, d := range dice {
		value += d.Value
	}
	return
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	opts := &Options{
		Rolls:   3,
		NumDice: 3,
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
		fmt.Printf("roll %d: %s (%d face value)\n", roll, dice, dice.FaceValue())

	}
}
