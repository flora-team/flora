package cmd

import (
	"fmt"
	"os"

	"github.com/Flora-team/flora-cli/model"

	"github.com/spf13/cobra"
)

func init() {
	newCmd.AddCommand(elementCmd)
}

var elementCmd = &cobra.Command{
	Use:   "element",
	Short: "Create a new element",
	Long:  `Create a new element`,
	Run: func(cmd *cobra.Command, args []string) {
		var savePath string
		if len(args) < 1 {
			savePath, _ = os.Getwd()
		} else {
			savePath = args[0]
		}
		createNewElement(savePath)
	},
}

func createNewElement(savePath string) {
	fmt.Println("This utility will walk you through creating a new element.")
	fmt.Println("Press ^C at any time to quit.")
	elementName := "exampleInput"
	fmt.Printf("element name: (%s) ", elementName)
	fmt.Scanf("%s\n", &elementName)
	elementDetails := "a simple Input, need 2 params to locate it."
	fmt.Printf("details: (%s) ", elementDetails)
	fmt.Scanf("%s\n", &elementDetails)

	elementFile := model.Element{
		ElementName:    elementName,
		ElementDetails: elementDetails,
		LocateParams:   []model.Param{{Param: "name", Comment: "the [name] attribute of input", Type: "string"}, {Param: "value", Comment: "the [value] attribute of input", Type: "string"}},
		LocatePattern: model.LocatePattern{
			Xpath: "//div[@class=\"cls\"]//input[@name=\"${name}\" and @value=\"${value}\"]",
		},
		Functions: []model.Function{{
			Name: "inputValue",
			Params: []model.Param{{
				Param:   "inputString",
				Comment: "the string to input",
				Type:    "string",
			}},
			Comment:   "input a string to the element",
			Operation: "setValue",
		}},
	}
	if err := elementFile.Save(savePath); err != nil {
		fmt.Println("ERROR: " + err.Error())
	} else {
		fmt.Println("element create successfully")
		fmt.Println("open the element file to add locate-related informations")
	}
}
