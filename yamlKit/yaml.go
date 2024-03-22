package yamlKit

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Unmarshal 需要搭配 yaml Tag
func Unmarshal(content []byte, ptr interface{}) error {
	return yaml.Unmarshal(content, ptr)
}

func UnmarshalFromFile(path string, prt interface{}) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return Unmarshal(bytes, prt)
}
