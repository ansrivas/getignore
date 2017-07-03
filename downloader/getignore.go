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
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/parnurzeal/gorequest"
)

// GetIgnore encapsulates request pointer and a languageMap which associates a
// integer string to a language. i.e. [1] -> "python"
type GetIgnore struct {
	request     *gorequest.SuperAgent
	languageMap map[string]string
}

// NewGitIgnore returns a new GetIgnore instance which can be used for listing
// and downloading, available .gitignore files.
func NewGitIgnore() *GetIgnore {
	request := gorequest.New()
	downloader := GetIgnore{
		request:     request,
		languageMap: make(map[string]string),
	}
	return &downloader
}

// ListLanguages displays the list of languages for which gitignore is available
func (gi *GetIgnore) ListLanguages(display bool) {
	var listResp response
	_, body, errs := gi.request.Get(baseIgnoreURL).End()
	check(errs)

	err := json.Unmarshal([]byte(body), &listResp)
	if err != nil {
		log.Fatal(err)
	}
	for idx, val := range listResp.Tree {
		lang := strings.Split(val.Path, ".gitignore")[0]
		if !strings.HasPrefix(lang, ".") {
			gi.languageMap[strings.ToLower(lang)] = lang
			if display == true {
				fmt.Println(idx, ":", lang)
			}
		}
	}
}

// DownloadFile downloads a gitignore for a given `language`
func (gi *GetIgnore) DownloadFile(language string) {

	//Populate the list
	gi.ListLanguages(false)

	//fetch the correct language
	language, err := normalizeString(language)
	if err != nil {
		log.Fatalf("\033[31mUnable to normalize %s to download\033[39m", language)
	}

	lang, ok := gi.languageMap[language]
	if !ok {
		log.Fatalf("\033[31mUnable to find gitignore for language `%s`.\033[39m\n\n", language)
	}

	fetchURL := parseLangURL(lang)

	resp, body, errs := gi.request.Get(fetchURL).End()
	if errs != nil {
		log.Fatalln("\033[31mFailed to download, please retry.\033[39m\n\n")
	}
	if resp.StatusCode == http.StatusNotFound {
		log.Fatalf("\033[31mWrong url queried: %s\033[39m\n\n", fetchURL)
	}
	writeFile(body, gitignoreFile)
}

func check(e []error) {
	if e != nil {
		log.Fatalln(e)
	}
}

// parseLangURL parses a given language and generates a proper link to download from github.
// eg. "https://raw.githubusercontent.com/github/gitignore/master/Python.gitignore"
func parseLangURL(lang string) string {
	u, err := url.Parse(lang + ".gitignore")
	if err != nil {
		log.Fatalln(err.Error())
	}
	base, err := url.Parse(downloadIgnoreURL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return base.ResolveReference(u).String()
}
