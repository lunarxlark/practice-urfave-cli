package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

var commands = []*cli.Command{
	{
		Name:      "foo",
		Aliases:   []string{"f"},
		Usage:     "usage of foo",
		UsageText: "what is different between Usage and UsageText",
	},
	{
		Name:      "bar",
		Aliases:   []string{"b"},
		Usage:     "usage of bar",
		UsageText: "what is different between Usage and UsageText",
	},
}

var exitCode int

func main() {
	if err := newApp().Run(os.Args); err != nil {
		fmt.Println(err.Error())
		exitCode = 1
	}
	os.Exit(exitCode)
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "practice"
	app.Usage = "test of cli building by urfave/cli"
	app.Version = "0.0.1" // I love semantic version
	app.Authors = []*cli.Author{
		{
			Name:  "foobar",
			Email: "foobar@gmail.com",
		},
		{
			Name:  "hogefuga",
			Email: "hogefuga@gmail.com",
		},
	}
	app.Description = "description of this command"
	app.UsageText = "what is different between Usage and UsageText?"
	app.Before = before
	app.After = after
	app.Commands = commands
	return app
}

func before(ctx *cli.Context) error {
	fmt.Println("starting...")
	return nil
}

func after(ctx *cli.Context) error {
	fmt.Println("finishing...")
	return nil
}
