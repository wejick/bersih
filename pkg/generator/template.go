package generator

import (
	"os"
	"text/template"
)

//Generate file from template
func Generate(filename string, templateString string, data interface{}) (err error) {
	file, err := os.Create(filename)
	if err != nil {
		return
	}
	tmplt, err := template.New("filename").Parse(templateString)
	if err != nil {
		return
	}
	err = tmplt.Execute(file, data)
	if err != nil {
		return
	}

	return
}
