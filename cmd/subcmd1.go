package cmd

import "github.com/urfave/cli/v2"

func cmdSubcmd1() *cli.Command {
	return &cli.Command{
		Name:   "subcmd1",
		Action: subcmd1,
		Usage:  "Uuage for subcmd1",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "subcmd1-arg1",
				Aliases: []string{"a"},
				Usage:   "usage for subcmd1-arg1",
			},
			&cli.BoolFlag{
				Name:  "subcmd1-arg2",
				Usage: "usage for subcmd1-arg2",
			},
		},
	}
}

func subcmd1(ctx *cli.Context) error {
	// get subcmd1 args
	// do something
	return nil
}
