package downloader

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/parnurzeal/gorequest"
)

var (
	baseURL     = "https://api.github.com/repos/github/gitignore/git/trees/HEAD"
	downloadURL = "https://raw.githubusercontent.com/github/gitignore/master/"
)

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

// GetIgnore encapsulates request pointer and a languageMap which associates a
// integer string to a language. i.e. [1] -> "python"
type GetIgnore struct {
	request     *gorequest.SuperAgent
	languageMap map[string]string
}

// New returns a new GetIgnore instance which can be used for listing
// and downloading, available .gitignore files.
func New() *GetIgnore {
	request := gorequest.New()
	downloader := GetIgnore{request: request,
		languageMap: make(map[string]string),
	}
	return &downloader
}

//ListLanguages displays the list of languages for which gitignore is available
func (gi *GetIgnore) ListLanguages(display bool) {
	var listResp response
	_, body, errs := gi.request.Get(baseURL).End()
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

// writeFile writes a given content to a file
func writeFile(content, lang string) {

	f, err := os.Create(".gitignore")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Printf("\033[32mSuccessfully downloaded .gitignore for %v in current working directory.\033[39m\n\n", lang)
	f.Sync()
}

// DownloadFile downloads a gitignore for a given `language`
func (gi *GetIgnore) DownloadFile(language string) {

	//Populate the list
	gi.ListLanguages(false)

	//fetch the correct language
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
	writeFile(body, lang)
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
	base, err := url.Parse(downloadURL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return base.ResolveReference(u).String()
}
