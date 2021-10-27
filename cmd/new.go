package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newCmd)
}

var newCmd = &cobra.Command{
	Use:   "new [page|element]",
	Short: "Create a new page or element",
	Long:  `Create a new page or element`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
