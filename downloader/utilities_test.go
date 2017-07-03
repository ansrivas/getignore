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
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_StripChars(t *testing.T) {
	assert := assert.New(t)
	actual := []string{"bsd-2-clause", "lgpl-3.0", "agpl-3-.0", "MIT-30.0", "Python"}
	expected := []string{"bsd2clause", "lgpl30", "agpl30", "mit300", "python"}

	for i, val := range actual {
		output, err := normalizeString(val)
		assert.Nil(err, "Should successfully replace")
		assert.Equal(expected[i], output, "Regex should have cleaned up the string properly")

	}
}
