package cmd

import (
	"fmt"
	"os"

	"github.com/Flora-team/flora/model"

	"github.com/spf13/cobra"
)

func init() {
	newCmd.AddCommand(pageCmd)
}

var pageCmd = &cobra.Command{
	Use:   "page",
	Short: "Create a new page",
	Long:  `Create a new page`,
	Run: func(cmd *cobra.Command, args []string) {
		var savePath string
		if len(args) < 1 {
			savePath, _ = os.Getwd()
		} else {
			savePath = args[0]
		}
		createNewPage(savePath)
	},
}

func createNewPage(savePath string) {
	fmt.Println("This utility will walk you through creating a new page.")
	fmt.Println("Press ^C at any time to quit.")
	pageName := "defaultPage"
	fmt.Printf("page name: (%s) ", pageName)
	fmt.Scanf("%s\n", &pageName)
	pageDetails := "a simple page"
	fmt.Printf("details: (%s) ", pageDetails)
	fmt.Scanf("%s\n", &pageDetails)

	pageFile := model.Page{
		PageName:    pageName,
		PageDetails: pageDetails,
	}
	if err := pageFile.Save(savePath); err != nil {
		fmt.Println("ERROR: " + err.Error())
	} else {
		fmt.Println("page create successfully")
	}

}
