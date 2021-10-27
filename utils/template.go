package utils

import (
	"bytes"
	"path/filepath"
	"text/template"
)

func Repeat(index int, length int, pattern string) string {
	if index < length-1 {
		return pattern
	} else {
		return ""
	}
}

func Add(a int, b int) int {
	return a + b
}

func ProcessFile(fileName string, vars interface{}) string {
	var err error
	tmpl := template.New(filepath.Base(fileName)).Funcs(template.FuncMap{"repeat": Repeat, "add": Add})
	tmpl, err = tmpl.ParseFiles(fileName)

	if err != nil {
		panic(err)
	}

	return process(tmpl, vars)
}

func ProcessTemplate(tmplVar string, vars interface{}) string {
	var err error
	tmpl := template.New(tmplVar).Funcs(template.FuncMap{"repeat": Repeat, "add": Add})
	tmpl, err = tmpl.Parse(tmplVar)

	if err != nil {
		panic(err)
	}

	return process(tmpl, vars)
}

func process(t *template.Template, vars interface{}) string {
	var tmplBytes bytes.Buffer

	err := t.Execute(&tmplBytes, vars)
	if err != nil {
		panic(err)
	}
	return tmplBytes.String()
}
