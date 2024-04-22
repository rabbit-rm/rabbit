package yamlToolkit

import (
	"fmt"
	"log"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	var cfg = `
password: "123456"
server:
    name: NameService
    port: 80
username: rabbit.rm`
	type server struct {
		Name string `yaml:"name"`
		Port int    `yaml:"port"`
	}

	type c struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Server   server `yaml:"server"`
	}
	var cc = &c{}
	err := Unmarshal([]byte(cfg), cc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*cc)
}
