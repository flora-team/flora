package core

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Flora-team/flora-cli/model"
	"github.com/Flora-team/flora-cli/templates"
	"github.com/Flora-team/flora-cli/utils"

	"gopkg.in/yaml.v2"
)

type codeSource struct {
	Page        model.Page
	Elements    []model.Element
	PackageName string
}

func StartGenerateCode(language string, sourcePath string, savePath string, packageName string, operationTransfer []string) {
	opt := make(map[string]string)
	for _, v := range operationTransfer {
		t := strings.Split(v, "=")
		opt[t[0]] = t[1]
	}

	cs := GenerateCodeSource(sourcePath, language, opt)

	fPath := savePath
	for _, v := range strings.Split(packageName, ".") {
		fPath = filepath.Join(fPath, v)
	}
	os.MkdirAll(fPath, os.ModePerm)
	pagesPath := filepath.Join(fPath, "pages")
	os.MkdirAll(pagesPath, os.ModePerm)
	modelsPath := filepath.Join(fPath, "models")
	os.MkdirAll(modelsPath, os.ModePerm)

	if language == "java" {
		// ===============================================
		s := GenerateElementCode(language, packageName)
		f, err := os.OpenFile(filepath.Join(modelsPath, "Element.java"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
		if err != nil {
			fmt.Printf("open err%s", err)
			return
		}
		defer f.Close()
		f.WriteString(s)
		// ===============================================
		// ===============================================
		s2 := GenerateLocatePatternCode(language, packageName)
		f2, err := os.OpenFile(filepath.Join(modelsPath, "LocatePattern.java"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
		if err != nil {
			fmt.Printf("open err%s", err)
			return
		}
		defer f2.Close()
		f2.WriteString(s2)
		// ===============================================
		// ===============================================
		s3 := GenerateBasePageCode(language, packageName)
		f3, err := os.OpenFile(filepath.Join(modelsPath, "BasePage.java"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
		if err != nil {
			fmt.Printf("open err%s", err)
			return
		}
		defer f3.Close()
		f3.WriteString(s3)
		// ===============================================

		for _, c := range cs {
			c.PackageName = packageName
			s := GeneratePageCode(language, c)
			f, err := os.OpenFile(filepath.Join(pagesPath, c.Page.PageName+".java"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
			if err != nil {
				fmt.Printf("open err%s", err)
				return
			}
			defer f.Close()
			f.WriteString(s)
		}
	} else if language == "robot" {
		// ===============================================
		s := GenerateRobotBaseCode(language, packageName)
		f, err := os.OpenFile(filepath.Join(pagesPath, "base.robot"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
		if err != nil {
			fmt.Printf("open err%s", err)
			return
		}
		defer f.Close()
		f.WriteString(s)
		// ===============================================
		for _, c := range cs {
			c.PackageName = packageName
			s := GeneratePageCode(language, c)
			f, err := os.OpenFile(filepath.Join(pagesPath, c.Page.PageName+".robot"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
			if err != nil {
				fmt.Printf("open err%s", err)
				return
			}
			defer f.Close()
			f.WriteString(s)
		}
	}

}
func GenerateCodeSource(path string, language string, opt map[string]string) []codeSource {
	fileInfoList, _ := ioutil.ReadDir(path)
	var codeSources []codeSource
	for i := range fileInfoList {
		if fileInfoList[i].IsDir() {
			var codeSource codeSource
			dirName := fileInfoList[i].Name()
			filepath.WalkDir(filepath.Join(path, dirName), func(path string, d fs.DirEntry, err error) error {
				files, _ := filepath.Glob(filepath.Join(path, "*.yaml"))
				for _, fileName := range files {
					if filepath.Base(fileName) == dirName+".yaml" {
						page, _ := ReadPage(fileName)
						codeSource.Page = page
					} else {
						element, err := ReadElement(fileName)
						if err != nil {
							fmt.Println("error in parse " + fileName)
							panic(err)
						}
						if language == "java" {
							element.LocatePattern.Xpath = strconv.Quote(element.LocatePattern.Xpath)
						}
						for j, function := range element.Functions {
							if operation, ok := opt[function.Operation]; ok {
								element.Functions[j].Operation = operation
							}
						}
						codeSource.Elements = append(codeSource.Elements, element)
					}
				}
				return nil
			})
			codeSources = append(codeSources, codeSource)

		} else {
			continue
		}
	}
	return codeSources
}

func ReadPage(path string) (model.Page, error) {
	page := new(model.Page)
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return *page, err
	}
	err = yaml.Unmarshal(yamlFile, page)
	if err != nil {
		return *page, err
	}
	return *page, nil
}

func ReadElement(path string) (model.Element, error) {
	element := new(model.Element)
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return *element, err
	}
	err = yaml.Unmarshal(yamlFile, element)
	if err != nil {
		return *element, err
	}
	return *element, nil
}

func GeneratePageCode(language string, c codeSource) string {
	if language == "java" {
		for i, element := range c.Elements {
			for j, function := range element.Functions {
				for k, param := range function.Params {
					if param.Type == "string" {
						c.Elements[i].Functions[j].Params[k].Type = "String"
					}
				}
			}
		}
	}
	return utils.ProcessTemplate(templates.PageTmpl(language), c)
}

func GenerateBasePageCode(language string, packageName string) string {
	return utils.ProcessTemplate(templates.BasePageTmpl(language), packageName)
}

func GenerateElementCode(language string, packageName string) string {
	return utils.ProcessTemplate(templates.ElementTmpl(language), packageName)
}

func GenerateLocatePatternCode(language string, packageName string) string {
	return utils.ProcessTemplate(templates.LocatePatternTmpl(language), packageName)
}

func GenerateRobotBaseCode(language string, packageName string) string {
	return utils.ProcessTemplate(templates.BaseTmpl(language), packageName)
}
