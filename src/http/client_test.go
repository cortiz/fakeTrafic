package http

import (
	"github.com/cortiz/fakeTrafic/src/yaml"
	"testing"
)

func TestDoRequest(t *testing.T) {
	request := yaml.Request{}
	request.ReadFile("../../test/echo-request.yaml")
	DoRequest(&request)

}
