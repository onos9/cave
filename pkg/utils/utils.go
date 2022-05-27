package utils

import (
	"bytes"
	"errors"
	"html/template"
	"io/ioutil"
	"os"
)

func ParseHtml(f string) (string, error) {
	bs, err := ioutil.ReadFile(f)
	if err != nil {
		return "", err
	}

	return string(bs), nil
}

func ParseTemplate(m interface{}) (string, error) {

	data := m.(map[string]interface{})
	dir, err := os.Getwd()
	if err != nil {
		return "", errors.New("ParseTemplate(): can't do os.Getwd")
	}

	filePath := dir + "/templates/" + data["filename"].(string)
	t, err := template.ParseFiles(filePath)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, m); err != nil {
		return "", err
	}

	return buf.String(), nil
}
