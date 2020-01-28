package main

import (
	"errors"
	"fmt"

	"github.com/urfave/cli/v2"
)

var commands = []*cli.Command{
	cmdFoo,
	cmdBar,
}

var cmdFoo = &cli.Command{
	Name:      "foo",
	Aliases:   []string{"f"},
	Usage:     "usage of foo",
	UsageText: "what is different between Usage and UsageText",
	Action:    doFoo,
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name: "x",
		},
	},
}

var cmdBar = &cli.Command{
	Name:      "bar",
	Aliases:   []string{"b"},
	Usage:     "usage of bar",
	UsageText: "what is different between Usage and UsageText",
	Action:    doBar,
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name: "y",
		},
	},
}

func doFoo(ctx *cli.Context) error {
	if ctx.Int("x") == 0 {
		return errors.New("x is not specified")
	}
	x := ctx.Int("x")
	for i := 1; i <= 10; i++ {
		switch i % x {
		case 0:
			fmt.Println("foo")
		default:
			fmt.Println(i)
		}
	}
	return nil
}

func doBar(ctx *cli.Context) error {
	fmt.Println("bar")
	return nil
}
