# getignore
Get your gitignore file directly from github

---
### Usage
```
$ getignore
Usage:
  getignore [command]

Available Commands:
  download    Download a gitignore file for the given language.
  list        Display a list of available gitignore files.

Use "getignore [command] --help" for more information about a command.
```

---
For convenience a binary is also committed in the github repo, under `prebuilt` directory, compiled using:

```go
GOOS=linux go build -ldflags="-s -w" github.com/ansrivas/getignore
```
