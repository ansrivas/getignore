// MIT License
//
// Copyright (c) 2017 Ankur Srivastava
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package downloader

import (
	"io/ioutil"
	"os"
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

	writeFile("Hello world", "testfile")
	defer func() {
		err := os.Remove("testfile")
		if err != nil {
			assert.Fail("Unable to write file")
		}
	}()

	dat, err := ioutil.ReadFile("testfile")

	assert.Nil(err, "WriteFile should have created `testfile` in pwd")
	assert.Equal(expected, string(dat), "Successfully write a file")

}
