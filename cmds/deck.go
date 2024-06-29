package cmds

import (
	"github.com/gabekus/kana/deck"
	cli "github.com/urfave/cli/v2"
	// "log"
)

func Deck() *cli.Command {
	return &cli.Command{
		Name:    "deck",
		Aliases: []string{"d"},
		Usage:   "Manage decks",
		Action: func(c *cli.Context) error {
			return PickDeck()
		},
		// Subcommands: []*cli.Command{
		// 	{
		// 		Name:    "add",
		// 		Aliases: []string{"a", "new"},
		// 		Usage:   "Add a deck",
		// 		Flags: []cli.Flag{
		// 			&cli.StringFlag{
		// 				Name:    "name",
		// 				Aliases: []string{"n"},
		// 				Usage:   "Deck name",
		// 			},
		// 		},
		// 		Action: func(c *cli.Context) error {
		// 			return deck.AddFromInteractivePrompt()
		// 		},
		// 	},
		// 	{
		// 		Name:    "list",
		// 		Aliases: []string{"l"},
		// 		Usage:   "List decks",
		// 		Action: func(c *cli.Context) error {
		// 			return deck.List()
		// 		},
		// 	},
		// },
	}
}

func PickDeck() error {
	return deck.Pick()
}
