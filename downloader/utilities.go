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
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var (
	baseIgnoreURL     = "https://api.github.com/repos/github/gitignore/git/trees/HEAD"
	downloadIgnoreURL = "https://raw.githubusercontent.com/github/gitignore/master/"
	baseLicensesURL   = "https://api.github.com/licenses"
	gitignoreFile     = ".gitignore"
	licenseFile       = "LICENSE"
)

// language struct is to unmarshal the response from baseURL which contains all
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

type licenseResponse struct {
	Licenses []license
}

type license struct {
	Key      string `json:"key"`
	Name     string `json:"name"`
	Spdxid   string `json:"spdx_id"`
	URL      string `json:"url"`
	Featured bool   `json:"featured"`
}

type licenseBody struct {
	Key            string   `json:"key"`
	Name           string   `json:"name"`
	Spdxid         string   `json:"spdx_id"`
	URL            string   `json:"url"`
	Featured       bool     `json:"featured"`
	HTMLURL        string   `json:"html_url"`
	Description    string   `json:"description"`
	Implementation string   `json:"implementation"`
	Permissions    []string `json:"permissions"`
	Conditions     []string `json:"conditions"`
	Limitations    []string `json:"limitations"`
	Body           string   `json:"body"`
}

func normalizeString(input string) (string, error) {

	re := regexp.MustCompile(`[-.]`)
	output := re.ReplaceAllString(input, "")
	output = strings.ToLower(output)
	return output, nil
}

// writeFile writes a given content to a file
func writeFile(content, filename string) {

	f, err := os.Create(filename)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Printf("\033[32mSuccessfully downloaded %v in current working directory.\033[39m\n\n", filename)
	f.Sync()
}
