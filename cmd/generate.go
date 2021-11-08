package cmd

import (
	"fmt"

	"github.com/flora-team/flora/core"
	"github.com/flora-team/flora/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	source            string
	target            string
	language          string
	packageName       string
	operationTransfer []string
	cfgFile           string
)

func init() {

	generateCmd.Flags().StringVarP(&cfgFile, "config", "c", "", "config file")

	generateCmd.Flags().String("source", "", "Source directory to read from")

	generateCmd.Flags().String("language", "", "Language of code generated")

	generateCmd.Flags().String("package", "", "The package of code")

	generateCmd.Flags().String("target", ".", "Target directory for generating code")

	generateCmd.Flags().StringArray("operation", []string{}, "Transfer default operation act")

	initConfig()

	rootCmd.AddCommand(generateCmd)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".flora" (without extension).
		// 当前目录下寻找配置文件
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName("flora.config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		viper.BindPFlags(generateCmd.Flags())
		// viper.BindPFlag("source", generateCmd.Flags().Lookup("source"))
		// viper.BindPFlag("language", generateCmd.Flags().Lookup("language"))
		// viper.BindPFlag("package", generateCmd.Flags().Lookup("package"))
		// viper.BindPFlag("target", generateCmd.Flags().Lookup("target"))
		// viper.BindPFlag("operation", generateCmd.Flags().Lookup("operation"))

	} else {
		fmt.Println(err.Error())
	}
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate Code from source files",
	Long:  `Generate Code from source files`,
	Run: func(cmd *cobra.Command, args []string) {
		language = viper.Get("language").(string)
		source = viper.Get("source").(string)
		target = viper.Get("target").(string)
		packageName = viper.Get("package").(string)
		operationTransfer := []string{}
		for _, v := range viper.Get("operation").([]interface{}) {
			operationTransfer = append(operationTransfer, v.(string))
		}

		if !utils.IsContainString([]string{"java", "robot"}, language) {
			fmt.Printf("unsupportted language: %s\n", language)
			return
		}

		core.StartGenerateCode(language, source, target, packageName, operationTransfer)
	},
}
