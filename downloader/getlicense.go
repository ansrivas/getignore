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

	"github.com/parnurzeal/gorequest"
)

// GetLicense encapsulates request pointer and a languageMap which associates a
// integer string to a language. i.e. [1] -> "python"
type GetLicense struct {
	request    *gorequest.SuperAgent
	licenseMap map[string]string
}

// NewLicense returns a new GetLicense instance which can be used for listing
// and downloading, available .license files.
func NewLicense() *GetLicense {
	request := gorequest.New()
	downloader := GetLicense{
		request:    request,
		licenseMap: make(map[string]string),
	}
	return &downloader
}

// DownloadFile downloads a gitignore for a given `language`
func (gl *GetLicense) DownloadFile(licenseInput string) {

	//Populate the list
	gl.ListLicenses(false)

	lindex, err := normalizeString(licenseInput)
	if err != nil {
		log.Fatalln("Unable to convert input string to a license.")
	}
	//fetch the correct language
	url, ok := gl.licenseMap[lindex]
	if !ok {
		log.Fatalf("\033[31mUnable to find license file for `%s`.\033[39m\n\n", lindex)
	}

	resp, body, errs := gl.request.Get(url).Set("Accept", "application/vnd.github.drax-preview+json").End()
	if errs != nil {
		log.Fatalln("\033[31mFailed to download, please retry.\033[39m\n\n")
	}

	var l licenseBody
	err = json.Unmarshal([]byte(body), &l)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == http.StatusNotFound {
		log.Fatalf("\033[31mWrong url queried: %s\033[39m\n\n", url)
	}
	writeFile(l.Body, licenseFile)
}

// ListLicenses displays the list of available licenses from github api
func (gl *GetLicense) ListLicenses(display bool) {
	var licResp licenseResponse
	_, body, errs := gl.request.Get(baseLicensesURL).Set("Accept", "application/vnd.github.drax-preview+json").End()
	check(errs)

	err := json.Unmarshal([]byte(body), &licResp.Licenses)
	if err != nil {
		log.Fatal(err)
	}
	for idx, val := range licResp.Licenses {
		output, err := normalizeString(val.Key)
		if err != nil {
			log.Fatalln("Unable to parse string", val.Key)
		}
		gl.licenseMap[output] = val.URL
		if display == true {
			fmt.Println(idx, ":", val.Key)
		}
	}
}
