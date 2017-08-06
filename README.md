# getignore

Get your gitignore file directly from github

[![Build Status](https://travis-ci.org/ansrivas/getignore.svg?branch=master)](https://travis-ci.org/ansrivas/getignore)

---
### Usage
```
→ getignore 
Usage:
  getignore [command]

Available Commands:
  downloadIgnore  Download a gitignore file for the given language.
  downloadLicense Download a license file from github.
  listIgnores     Display a list of available gitignore files.
  listLicenses    Display a list of available licenses.

Use "getignore [command] --help" for more information about a command.
```

---
For convenience a binary is also committed in the github repo, under `prebuilt` directory, compiled using:

```go
GOOS=linux go build -ldflags="-s -w" github.com/ansrivas/getignore
```
