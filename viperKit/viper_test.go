package viperKit

import (
	"fmt"
	"log"
	"testing"
)

func TestRead(t *testing.T) {
	var cfg = `
password: "123456"
server:
    name: NameService
    port: 80
username: rabbit.rm`

	v, err := read([]byte(cfg), YAML, nil)
	if err != nil {
		log.Fatal(err)
	}
	for k, vv := range v.AllSettings() {
		fmt.Printf("k:%v,vv:%v\n", k, vv)
	}
}

func TestUnmarshal(t *testing.T) {
	var cfg = `
password: "123456"
server:
    name: NameService
    port: 80
username: rabbit.rm`

	type server struct {
		Name string `json:"name"`
		Port int    `json:"port"`
	}

	type c struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Server   server `json:"server"`
	}

	var cc = &c{}
	err := Unmarshal([]byte(cfg), YAML, nil, cc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*cc)
}
