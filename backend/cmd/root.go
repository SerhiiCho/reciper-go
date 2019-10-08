package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "reciper",
	Short: "Recipe web application CLI",
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Usage()
		fmt.Println(err)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
