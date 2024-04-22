package viperToolkit

type configType string

const (
	JSON       configType = "json"
	YAML       configType = "yaml"
	PROPERTIES configType = "properties"
	INI        configType = "ini"
)
