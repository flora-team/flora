package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// Used for flags.

	rootCmd = &cobra.Command{
		Use:   "flora",
		Short: "Web element manager for UI test",
		Long: `Flora is a tool to manage web elements for UI test.
By using flora, you can generate code in different languages from a unified configuration file.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {

}
