package yaml

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)


type Request struct {
	Method string `yaml:"method"`
	Url string `yaml:"url"`
	ContentType string `yaml:"content-type"`
	Headers map[string]string `yaml:"headers"`
	Data map[string]string `yaml:"data"`
}

func (r *Request) ReadFile(file string){
	yamlFile, err :=ioutil.ReadFile(file)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, r)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}
