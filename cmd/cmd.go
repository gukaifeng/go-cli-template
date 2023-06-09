package cmd

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func myprogram(c *cli.Context) error {
	// get global args
	// do something
	return nil
}

func globalFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  "arg1",
			Usage: "usage for arg1",
		},
		&cli.BoolFlag{
			Name:    "arg2",
			Aliases: []string{"a2"},
			Value:   false,
			Usage:   "usage for arg2",
		},
	}
}

func Main() {
	app := &cli.App{
		Name:   "myprogram",
		Usage:  "action with no subcommand",
		Action: myprogram, // action with no subcommand
		Flags:  globalFlags(),
		Commands: []*cli.Command{
			cmdSubcmd1(),
			cmdSubcmd2(),
		},
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%v run error: %v\n", app.Name, err)
	}
}
