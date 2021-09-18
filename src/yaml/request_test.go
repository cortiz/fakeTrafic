package yaml

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRequest_readFile(t *testing.T) {
	var req = Request{}
	req.ReadFile("../../test/sample-request.yaml")
	assert.Equal(t, "POST", req.Method)
	assert.Equal(t, "https://test.com", req.Url)
	assert.Equal(t, 2, len(req.Headers))
	assert.Equal(t, 2, len(req.Data))
}
