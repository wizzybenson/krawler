package contact

import (
	"github.com/spf13/cobra"
)

var ContactCmd = &cobra.Command{
    Use:  "contact",
    Short: "contact - object contact to attach actions like grabemails",
    
    Run: func(cmd *cobra.Command, args []string) {

    },
}

func init() {
    ContactCmd.AddCommand(grabEmailsCmd)
}