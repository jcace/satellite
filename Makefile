.DEFAULT_GOAL := build
GITIGNORE ?= go

PLATFORMS ?= linux/amd64 darwin/amd64 darwin/arm64
temp = $(subst /, ,$@)
os = $(word 1, $(temp))
arch = $(word 2, $(temp))

build: $(PLATFORMS)

$(PLATFORMS):
	GOOS=$(os) GOARCH=$(arch) go build -o ./artifacts/remote_desktop_ping_service_$(os)_$(arch) main.go

run:
	go run main.go -url ${URL}

test:
	go test ./...

gitignore:
	curl -Ls "http://www.gitignore.io/api/$(GITIGNORE)" | tee .gitignore
	@if [ -f .gitignore.custom ]; then \
		cat .gitignore.custom >> .gitignore; \
	fi

clean:
	rm -rf reservation_daemon_*

.PHONY: build clean run
