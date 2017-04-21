# getignore
Get your gitignore file directly from github

---
For convenience a binary is also committed in the github repo, under `prebuilt` directory, compiled using:

```go
GOOS=linux go build -ldflags="-s -w" github.com/ansrivas/getignore
```
