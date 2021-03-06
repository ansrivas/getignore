.DEFAULT_GOAL := help

help:          ## Show available options with this Makefile
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

.PHONY : test
test:          ## Run all the tests
test:
	@GOOS=linux go build -ldflags="-s -w" github.com/ansrivas/getignore && \
	./test.sh

clean:         ## Clean the application
clean:
	@go clean -i ./...
	@rm -rf ./getignore

release:       ## Create a release build
release:	clean
	@GOOS=linux go build -ldflags="-s -w" github.com/ansrivas/getignore
