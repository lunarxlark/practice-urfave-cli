package main

import (
	"errors"
	"fmt"

	"github.com/urfave/cli/v2"
)

var commands = []*cli.Command{
	{
		Name:      "foo",
		Aliases:   []string{"f"},
		Usage:     "usage of foo",
		UsageText: "what is different between Usage and UsageText",
		Action:    cmdFoo,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name: "x",
			},
		},
	},
	{
		Name:      "bar",
		Aliases:   []string{"b"},
		Usage:     "usage of bar",
		UsageText: "what is different between Usage and UsageText",
		Action:    cmdBar,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name: "y",
			},
		},
	},
}

func cmdFoo(ctx *cli.Context) error {
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

func cmdBar(ctx *cli.Context) error {
	fmt.Println("bar")
	return nil
}
