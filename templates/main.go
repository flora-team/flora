package templates

import (
	javaModels "github.com/Flora-team/flora-cli/templates/java/models"

	javaPages "github.com/Flora-team/flora-cli/templates/java/pages"
	robotPages "github.com/Flora-team/flora-cli/templates/robot/pages"
)

func BasePageTmpl(language string) string {
	if language == "java" {
		return javaModels.BasePage
	}
	return ""
}

func ElementTmpl(language string) string {
	if language == "java" {
		return javaModels.Element
	}
	return ""
}

func LocatePatternTmpl(language string) string {
	if language == "java" {
		return javaModels.LocatePattern
	}
	return ""
}

func PageTmpl(language string) string {
	if language == "java" {
		return javaPages.Page
	} else if language == "robot" {
		return robotPages.Page
	}
	return ""
}

func BaseTmpl(language string) string {
	if language == "robot" {
		return robotPages.Base
	}
	return ""
}
