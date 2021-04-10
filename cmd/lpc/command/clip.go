package command

import (
	"context"
	"fmt"
	"log"

	"github.com/ansd/lastpass-go"
	"github.com/atotto/clipboard"
	"github.com/sh3rp/lpc"
	"github.com/spf13/cobra"
)

var clipCmd = &cobra.Command{
	Use:   "clip",
	Short: "Sends secret to clipboard",
	Long:  "Retries a specified secret from LastPass and puts it into the clipboard for pasting",
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
			clipboard.WriteAll(acct.Password)
			fmt.Printf("%s written to clipboard.\n", args[0])
		} else {
			fmt.Printf("%s not found.\n", args[0])
		}
	},
}
