package contact

import (
	"github.com/spf13/cobra"
	"github.com/wizzybenson/krawler/pkg/contact"
)

var (
	Filename     = ""
	MaxLength    = 0
	CountryCode  = ""
	LanguageCode = ""
	Query        = ""
)

var grabEmailsCmd = &cobra.Command{
	Use:   "grabemails",
	Short: "Performs google search and returns contact-emails from the serps sites",
	Run: func(cmd *cobra.Command, args []string) {
		contact.GrabEmails(Filename, CountryCode, LanguageCode, Query, MaxLength)
	},
}

func init() {
	ContactCmd.AddCommand(grabEmailsCmd)
	grabEmailsCmd.Flags().StringVarP(&Filename, "filename", "f", "contact.csv", "Filename to save the csv")
	grabEmailsCmd.Flags().IntVarP(&MaxLength, "max-length", "m", 1000, "Max length of the results returned by google")
	grabEmailsCmd.Flags().StringVarP(&CountryCode, "country-code", "c", "ch", "Country code to use for the google search")
	grabEmailsCmd.Flags().StringVarP(&LanguageCode, "lang", "l", "fr", "Lanuguage code to use for the google search")
	grabEmailsCmd.Flags().StringVarP(&Query, "query", "q", "Centre esthétique", "Query to search in google")
}
