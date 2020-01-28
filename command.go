package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var commands = []*cli.Command{
	{
		Name:      "foo",
		Aliases:   []string{"f"},
		Usage:     "usage of foo",
		UsageText: "what is different between Usage and UsageText",
		Action:    foo,
	},
	{
		Name:      "bar",
		Aliases:   []string{"b"},
		Usage:     "usage of bar",
		UsageText: "what is different between Usage and UsageText",
		Action:    bar,
	},
}

func foo(ctx *cli.Context) error {
	fmt.Println("foo")
	return nil
}

func bar(ctx *cli.Context) error {
	fmt.Println("bar")
	return nil
}
