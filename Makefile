LOCAL_UID=$(shell id -u)
LOCAL_GID=$(shell id -g)
OUTPUT_BINARY=bin/inc-zone-soa
PACKAGE=allypost.net/binder
define LDFLAGS
-X '$(PACKAGE)/app/version.buildTimestamp=$(shell date -u '+%Y-%m-%dT%H:%M:%S%z')'
-X '$(PACKAGE)/app/version.buildProgramName="$(shell basename "$(OUTPUT_BINARY)")"'
endef
LDFLAGS:=$(strip $(LDFLAGS))

.PHONY: all
all: $(OUTPUT_BINARY)

.PHONY: clean
clean:
	rm -f $(OUTPUT_BINARY)

$(OUTPUT_BINARY):
	$(MAKE) build

.PHONY: run
run: build
	@$(OUTPUT_BINARY) ./dev-test/* || exit 0

.PHONY: build
build:
	CGO_ENABLED=0 \
	go \
	build \
	-a \
	-tags osusergo,netgo \
	-gcflags "all=-N -l" \
	-ldflags="-s -w -extldflags \"-static\" $(LDFLAGS)" \
	-o "${OUTPUT_BINARY}" \
	main.go

.PHONY: format
format:
	gofmt -e -l -s -w .

.PHONY: fmt
fmt: format

.PHONY: sync-deps
sync-deps:
	CGO_ENABLED=0 go mod download