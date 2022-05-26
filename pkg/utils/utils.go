package utils

import (
	"bytes"
	"errors"
	"html/template"
	"io/ioutil"
	"path"
	"runtime"
)

func ParseHtml(f string) (string, error) {
	bs, err := ioutil.ReadFile(f)
	if err != nil {
		return "", err
	}

	return string(bs), nil
}

func ParseTemplate(m interface{}) (string, error) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return "", errors.New("can not get filename")
	}
	
	data := m.(map[string]interface{})
	dir := path.Dir(filename)
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
