package cmd

import "github.com/urfave/cli/v2"

func cmdSubcmd2() *cli.Command {
	return &cli.Command{
		Name:   "subcmd2",
		Action: subcmd2,
		Usage:  "Uuage for subcmd2",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "subcmd2-arg1",
				Aliases: []string{"a"},
				Usage:   "usage for subcmd2-arg1",
			},
			&cli.BoolFlag{
				Name:  "subcmd2-arg2",
				Usage: "usage for subcmd2-arg2",
			},
		},
	}
}

func subcmd2(ctx *cli.Context) error {
	// get subcmd2 args
	// do something
	return nil
}
