package model

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/Flora-team/flora-cli/utils"

	"gopkg.in/yaml.v2"
)

type Page struct {
	PageName    string `yaml:"pageName"`
	PageDetails string `yaml:"pageDetails"`
}

func (p *Page) Save(savePath string) error {
	pagePath := path.Join(savePath, p.PageName)
	fileName := path.Join(pagePath, p.PageName+".yaml")
	if utils.IsPathExist(pagePath) {
		return fmt.Errorf("directory %s exists", pagePath)
	}
	os.MkdirAll(pagePath, os.ModePerm)
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("error opening/creating file %s", fileName)
	}
	defer file.Close()

	enc := yaml.NewEncoder(file)

	err = enc.Encode(p)
	if err != nil {
		return errors.New("error encoding")
	}

	return nil
}
