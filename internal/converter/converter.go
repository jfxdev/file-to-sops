package converter

import (
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v3"
)

const artifactFile string = "file.yaml"

type Data struct {
	Metadata Metadata `yaml:"metadata"`
	Content  string   `yaml:"content"`
}

type Metadata struct {
	FileName string `yaml:"file_name"`
}

func Parse(fileName string) error {
	if fileName == "" {
		return fmt.Errorf("a file path must be specified")
	}

	body, err := ioutil.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("unable to read file: %v", err.Error())
	}

	encoded := b64.StdEncoding.EncodeToString(body)

	data := Data{
		Metadata: Metadata{
			FileName: fileName,
		},
		Content: encoded,
	}

	content, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	list := strings.Split(fileName, "/")
	if len(list) > 0 {
		list = list[:len(list)-1]
		path := strings.Join(list, "/")
		err = ioutil.WriteFile(fmt.Sprintf("%s/%s", path, artifactFile), content, 0644)
		if err != nil {
			return err
		}
	}

	return nil

}
