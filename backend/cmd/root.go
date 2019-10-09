package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "reciper",
	Short: "Recipe web application CLI",
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Usage()

		if err != nil {
			log.Print(err)
		}
	},
}

// Execute cobra smd command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
