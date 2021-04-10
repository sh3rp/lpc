package command

import (
	"errors"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var username = os.Getenv("LP_USER")
var password = os.Getenv("LP_PASS")

var trusted bool

var rootCmd = &cobra.Command{
	Use:   "lpc",
	Short: "LPC v1.0",
	Long:  "LastPass Client - v1.0",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var CheckSecretName = func(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("Need to specify a secret name")
	}
	return nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&trusted, "trusted", "t", false, "Cache 2FA authentication")
	rootCmd.AddCommand(clipCmd)
	rootCmd.AddCommand(outCmd)
}
