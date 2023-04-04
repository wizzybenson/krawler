package krawler

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
    "github.com/wizzybenson/krawler/cmd/contact"
)
var version = "0.0.1"
var KrawlerCmd = &cobra.Command{
    Use:  "krawler",
    Short: "krawler - a cli tool to perform various web crawling needs",
    Version: version,
    Run: func(cmd *cobra.Command, args []string) {

    },
}

func Execute() {
    if err := KrawlerCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
        os.Exit(1)
    }
}

func init() {
    KrawlerCmd.AddCommand(contact.ContactCmd)
}

