package viperKit

import (
	"bytes"

	"github.com/spf13/viper"
)

func read(content []byte, cfgType configType, defaultMap map[string]interface{}) (*viper.Viper, error) {
	v := viper.New()
	for k, vv := range defaultMap {
		v.SetDefault(k, vv)
	}
	v.SetConfigType(string(cfgType))
	if err := v.ReadConfig(bytes.NewBuffer(content)); err != nil {
		return nil, err
	}
	return v, nil
}

func readFile(path string, defaultMap map[string]interface{}) (*viper.Viper, error) {
	v := viper.New()
	for k, vv := range defaultMap {
		v.SetDefault(k, vv)
	}
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	return v, nil
}
