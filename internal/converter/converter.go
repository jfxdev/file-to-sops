package converter

import (
	"bytes"
	"crypto/sha512"
	b64 "encoding/base64"
	"fmt"
	"io"
	"io/fs"
	"os"

	"gopkg.in/yaml.v3"
)

const artifactFile string = "file.yaml"

type Data struct {
	Metadata Metadata `yaml:"metadata"`
	Content  string   `yaml:"content"`
}

type Metadata struct {
	FileName string `yaml:"file_name"`
	CheckSum string `yaml:"check_sum"`
}

func ToArtifact(fileName string) (err error) {
	content, sum, err := readFile(fileName)
	if err != nil {
		return
	}

	encodedContent := b64.StdEncoding.EncodeToString(content)

	data := Data{
		Content: encodedContent,
		Metadata: Metadata{
			FileName: fileName,
			CheckSum: sum,
		},
	}

	content, err = yaml.Marshal(data)
	if err != nil {
		return err
	}

	err = writeFile(artifactFile, content, 0644)
	return
}

func ToFile(artifact string) (err error) {
	var data Data
	content, _, err := readFile(artifact)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(content, &data)
	if err != nil {
		return err
	}

	decodedContent, err := b64.StdEncoding.DecodeString(data.Content)
	if err != nil {
		return err
	}

	sum, err := calcSum(decodedContent)
	if err != nil {
		return err
	}

	if sum != data.Metadata.CheckSum {
		return fmt.Errorf("checksum metadata does not match with file content")
	}

	err = writeFile(data.Metadata.FileName, decodedContent, 0644)
	return
}

func readFile(filePath string) (content []byte, sum string, err error) {
	if filePath == "" {
		return content, sum, fmt.Errorf("a file path must be specified")
	}

	content, err = os.ReadFile(filePath)
	if err != nil {
		return content, sum, fmt.Errorf("unable to read file: %v", err.Error())
	}

	sum, err = calcSum(content)
	return
}

func writeFile(filePath string, content []byte, perm fs.FileMode) error {
	return os.WriteFile(filePath, content, 0644)
}

func calcSum(content []byte) (sum string, err error) {
	hash := sha512.New()
	_, err = io.Copy(hash, bytes.NewReader(content))
	if err != nil {
		return
	}
	sum = fmt.Sprintf("%x", hash.Sum(nil))
	return
}
