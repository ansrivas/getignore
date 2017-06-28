package downloader

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParseLangUrl(t *testing.T) {
	assert := assert.New(t)
	expected := "https://raw.githubusercontent.com/github/gitignore/master/Python.gitignore"
	actual := parseLangURL("Python")
	assert.Equal(expected, actual, "Urls must be formatted properly for downloads")
}

func Test_WriteFile(t *testing.T) {
	assert := assert.New(t)
	expected := "Hello world"

	writeFile("Hello world", "")
	dat, err := ioutil.ReadFile(".gitignore")
	assert.Nil(err, "WriteFile should have created `.gitignore` in pwd")
	assert.Equal(expected, string(dat), "Successfully write a file")
}
