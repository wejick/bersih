package generator

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	type DataTest struct {
		Data string
	}
	testTemplate := `{{.Data}} is data`
	err := Generate("testFile", testTemplate, DataTest{Data: "data"})
	assert.NoError(t, err, "Generate test failed")
	testFile, err := ioutil.ReadFile("testFile")
	assert.NoError(t, err, "Generate test failed : can't open output")
	assert.Equal(t, string(testFile[:]), "data is data")

	os.Remove("testFile")
}
