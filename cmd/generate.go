package cmd

import (
	"fmt"

	"github.com/Flora-team/flora/core"
	"github.com/Flora-team/flora/utils"

	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	source            string
	target            string
	language          string
	packageName       string
	operationTransfer []string
)

func init() {
	generateCmd.Flags().StringVarP(&source, "source", "s", "", "Source directory to read from")
	generateCmd.MarkFlagRequired("source")
	generateCmd.Flags().StringVarP(&language, "language", "l", "", "Language of code generated")
	generateCmd.MarkFlagRequired("language")
	generateCmd.Flags().StringVarP(&packageName, "package", "p", "", "The package of code")
	generateCmd.MarkFlagRequired("package")

	generateCmd.Flags().StringVarP(&target, "target", "t", ".", "Target directory for generating code")
	generateCmd.Flags().StringArrayVarP(&operationTransfer, "operation", "o", []string{}, "Transfer default operation act")

	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate Code from source files",
	Long:  `Generate Code from source files`,
	Run: func(cmd *cobra.Command, args []string) {
		if !utils.IsContainString([]string{"java", "robot"}, language) {
			fmt.Printf("unsupportted language: %s\n", language)
			return
		}
		core.StartGenerateCode(language, source, target, packageName, operationTransfer)
	},
}
