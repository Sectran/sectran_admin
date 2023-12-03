OS ?=$(shell go env GOOS)
ARCH ?=$(shell go env GOARCH)
COMMIT := $(shell git rev-parse --short HEAD)
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
LDFLGGS=""

.PHONY: build
build:
	@mkdir -p build
	go build -ldflags ${LDFLGGS} -o build/sectran-${OS}-${ARCH} cmd/sectran_admin/main.go 
	
.PHONY: debug 
debug: build

.PHONY: release 
release: build

.PHONY: clean
clean:
	rm -rf build/*