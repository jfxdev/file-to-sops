package app

import (
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"sops-for-files/app/config"
	"strings"

	"gopkg.in/yaml.v3"
)

const secretArtifactName string = "secret.yaml"

type Data struct {
	Metadata Metadata `yaml:"metadata"`
	Content  string   `yaml:"content"`
}

type Metadata struct {
	Path string `yaml:"path"`
}

func Single(filepath string) error {
	if filepath == "" {
		return fmt.Errorf("a file path must be specified")
	}

	body, err := ioutil.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("unable to read file: %v", err.Error())
	}

	if err := config.Get().Files.Valid(filepath); err != nil {
		return err
	}

	encoded := b64.StdEncoding.EncodeToString(body)

	data := Data{
		Metadata: Metadata{
			Path: filepath,
		},
		Content: encoded,
	}

	content, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	list := strings.Split(filepath, "/")
	if len(list) > 0 {
		list = list[:len(list)-1]
		path := strings.Join(list, "/")
		err = ioutil.WriteFile(fmt.Sprintf("%s/%s", path, secretArtifactName), content, 0644)
		if err != nil {
			return err
		}
	}

	return nil

}
