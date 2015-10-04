package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
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

	// keep rolling until these die appear
	RollUntil string
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

// return a number of dice rolled
func NewDice(n, sides int) Dice {
	dice := make(Dice, n)
	for i := 0; i < n; i++ {
		dice[i] = NewRoll(sides)
	}
	return dice
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
func (dice Dice) Values() []int {
	values := make([]int, 0)
	for _, d := range dice {
		values = append(values, d.Value)
	}
	return values

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
	flag.StringVar(&opts.RollUntil, "roll-until", opts.RollUntil, "roll until the sequence is met (coma delimited)")
	flag.Parse()

	if opts.RollUntil != "" {

		tmp := strings.Split(opts.RollUntil, ",")

		if opts.NumDice > len(tmp) {
			fmt.Printf("setting number of dice to %d\n", len(tmp))
			opts.NumDice = len(tmp)
		}

		roll_until := make([]int, 0)
		for _, v := range tmp {
			n, err := strconv.Atoi(v)
			if err != nil {
				fmt.Printf("ERROR: %s: %s\n", v, err)
				continue
			}
			roll_until = append(roll_until, n)
		}

		fmt.Printf("rolling until sequence %v\n", roll_until)

		var dice Dice

		var num_rolls int
		for num_rolls = 0; ; num_rolls++ {
			dice = NewDice(opts.NumDice, opts.Sides)
			dice.Roll()

			met := true
			for i, v := range dice.Values() {

				if roll_until[i] != v {
					met = false
					break
				}
			}

			if met == true {
				break
			}

		}
		fmt.Printf("rolled %s in %d rolls\n", tmp, num_rolls)

		return
	}

	for roll := 1; roll <= opts.Rolls; roll++ {
		dice := NewDice(opts.NumDice, opts.Sides)
		dice.Roll()
		fmt.Printf("roll %d: %s (%d face value)\n", roll, dice, dice.FaceValue())
	}

}
