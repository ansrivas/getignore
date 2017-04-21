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
	BaseURL     = "https://api.github.com/repos/github/gitignore/git/trees/HEAD"
	DownloadURL = "https://raw.githubusercontent.com/github/gitignore/master/"
)

type Language struct {
	Path string `json:"path"`
	Mode string `json:"mode"`
	Type string `json:"type"`
	Sha  string `json:"sha"`
	Size int    `json:"size"`
	Url  string `json:"url"`
}

type Response struct {
	Sha  string     `json:"sha"`
	Url  string     `json:"url"`
	Tree []Language `json:"tree"`
}

type Downloader struct {
	request     *gorequest.SuperAgent
	languageMap map[string]string
}

func New() *Downloader {
	request := gorequest.New()
	downloader := Downloader{request: request,
		languageMap: make(map[string]string),
	}
	return &downloader
}

//ListLanguages displays the list of languages for which gitignore is available
func (d *Downloader) ListLanguages(display bool) {
	var listResp Response
	_, body, errs := d.request.Get(BaseURL).End()
	check(errs)

	err := json.Unmarshal([]byte(body), &listResp)
	if err != nil {
		log.Fatal(err)
	}
	for idx, val := range listResp.Tree {
		lang := strings.Split(val.Path, ".gitignore")[0]
		if !strings.HasPrefix(lang, ".") {
			d.languageMap[strings.ToLower(lang)] = lang
			if display == true {
				fmt.Println(idx, ":", lang)
			}
		}
	}
}

//writeFile writes a given content to  a file
func WriteFile(content, lang string) {

	f, err := os.Create(".gitignore")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Printf("\033[32mSuccessfully downloaded .gitignore for %v in present working directory.\033[39m\n\n", lang)
	f.Sync()
}

//DownloadFile downloads a gitignore for a given `language`
func (d *Downloader) DownloadFile(language string) {

	//Populate the list
	d.ListLanguages(false)

	//fetch the correct language
	lang, ok := d.languageMap[language]
	if !ok {
		log.Fatalf("\033[31mUnable to find gitignore for language `%s`.\033[39m\n\n", language)
	}

	fetchURL := ParseLangUrl(lang)

	resp, body, errs := d.request.Get(fetchURL).End()
	if errs != nil {
		log.Fatalln("\033[31mFailed to download, please retry.\033[39m\n\n")
	}
	if resp.StatusCode == http.StatusNotFound {
		log.Fatalf("\033[31mWrong url queried: %s\033[39m\n\n", fetchURL)
	}
	WriteFile(body, lang)
}

func check(e []error) {
	if e != nil {
		log.Fatalln(e)
	}
}

//ParseLangUrl parses a given language and generates a proper link to download from github.
// eg. "https://raw.githubusercontent.com/github/gitignore/master/Python.gitignore"
func ParseLangUrl(lang string) string {
	u, err := url.Parse(lang + ".gitignore")
	if err != nil {
		log.Fatalln(err.Error())
	}
	base, err := url.Parse(DownloadURL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return base.ResolveReference(u).String()
}
