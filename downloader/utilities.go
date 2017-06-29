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

var (
	baseURL     = "https://api.github.com/repos/github/gitignore/git/trees/HEAD"
	downloadURL = "https://raw.githubusercontent.com/github/gitignore/master/"
)

// language struct is to unmarshal the respone from baseURL which contains all
// the languages with their corresponding url for .gitignore files.
type language struct {
	Path  string `json:"path"`
	Mode  string `json:"mode"`
	Stype string `json:"type"`
	Sha   string `json:"sha"`
	Size  int    `json:"size"`
	URL   string `json:"url"`
}

type response struct {
	Sha  string     `json:"sha"`
	URL  string     `json:"url"`
	Tree []language `json:"tree"`
}
