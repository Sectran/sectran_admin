version=0.1.0
revision=1

os ?=$(shell go env GOOS)
arch ?=$(shell go env GOARCH)

commit := $(shell git rev-parse --short HEAD)
branch := $(shell git rev-parse --abbrev-ref HEAD)
time := $(shell date +'%Y-%m-%dT%H:%M:%S%z')
ldflags := ""

.PHONY: init
init:
	@echo "Building with commit: $(commit), branch: $(branch), time: $(time)"
	@echo "generating version.go"
	@echo "package version" > cli/version/version.go
	@echo "" >> cli/version/version.go
	@echo "var (" >> cli/version/version.go
	@echo "    Commit string = \"$(commit)\"" >> cli/version/version.go
	@echo "    Branch string = \"$(branch)\"" >> cli/version/version.go
	@echo "    BuildTime string = \"$(time)\"" >> cli/version/version.go
	@echo ")" >> cli/version/version.go

.PHONY: build
build: init
	go build -ldflags ${ldflags} -o bin/sectran-${os}-${arch}
.PHONY: clean
clean:
	@if [ -e bin/sectran-${os}-${arch} ]; then rm -f bin/sectran-${os}-${arch}; fi
	@if [ -e cli/version/version.go ]; then rm -f cli/version/version.go; fi