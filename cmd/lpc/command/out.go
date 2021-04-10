package command

import (
	"context"
	"fmt"
	"log"

	"github.com/ansd/lastpass-go"
	"github.com/sh3rp/lpc"
	"github.com/spf13/cobra"
)

var outCmd = &cobra.Command{
	Use:   "out",
	Short: "Outputs the secret",
	Long:  "Outputs the specified secret to stdout (useful for referencing in scripts ala $(lpc out secret.name))",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		accounts, err := lpc.GetSecrets(context.Background(), username, password, trusted)

		if err != nil {
			log.Fatal(err)
		}

		var acct *lastpass.Account

		for _, account := range accounts {
			if account.Name == args[0] {
				acct = account
				break
			}
		}

		if acct != nil {
			fmt.Printf("%s", acct.Password)
		}
	},
}
