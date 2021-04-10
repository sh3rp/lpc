package command

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var username = os.Getenv("LP_USER")
var password = os.Getenv("LP_PASS")

var trusted bool

var rootCmd = &cobra.Command{
	Use: "lpc",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&trusted, "t", false, "Cache 2FA authentication")

	rootCmd.AddCommand(clipCmd)
	rootCmd.AddCommand(outCmd)
}
