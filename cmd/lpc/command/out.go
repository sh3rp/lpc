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
	Use: "out",
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