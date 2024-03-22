package viperKit

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

func Unmarshal(content []byte, cfgType configType, defaultMap map[string]interface{}, ptr interface{}) error {
	v, err := read(content, cfgType, defaultMap)
	if err != nil {
		return err
	}
	return unmarshal(v, ptr)
}
func UnmarshalFromFile(path string, defaultMap map[string]interface{}, ptr interface{}) error {
	v, err := readFile(path, defaultMap)
	if err != nil {
		return err
	}
	return unmarshal(v, ptr)
}

func unmarshal(v *viper.Viper, ptr interface{}) error {
	return v.Unmarshal(ptr, func(dc *mapstructure.DecoderConfig) {
		dc.Squash = true
	})
}
