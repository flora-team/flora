package model

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/Flora-team/flora/utils"

	"gopkg.in/yaml.v2"
)

type LocatePattern struct {
	Xpath string `yaml:"xpath"`
}

type Param struct {
	Param   string `yaml:"param"`
	Comment string `yaml:"comment"`
	Type    string `yaml:"type"`
}

type Function struct {
	Name      string  `yaml:"name"`
	Params    []Param `yaml:"params"`
	Comment   string  `yaml:"comment"`
	Operation string  `yaml:"operation"`
}

type Element struct {
	ElementName    string        `yaml:"elementName"`
	ElementDetails string        `yaml:"elementDetails"`
	IsBaseElement  bool          `yaml:"isBaseElement"`
	LocateParams   []Param       `yaml:"locateParams"`
	LocatePattern  LocatePattern `yaml:"locatePattern"`
	Functions      []Function    `yaml:"functions"`
}

func (e *Element) Save(savePath string) error {
	fileName := path.Join(savePath, e.ElementName+".yaml")
	if utils.IsPathExist(fileName) {
		return fmt.Errorf("file %s exists", fileName)
	}
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("error opening/creating file %s", fileName)
	}
	defer file.Close()

	enc := yaml.NewEncoder(file)

	err = enc.Encode(e)
	if err != nil {
		return errors.New("error encoding")
	}
	return nil
}
