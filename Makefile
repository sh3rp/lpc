
build:
	go build -o lpc cmd/lpc/lpc.go

install:
	go install cmd/lpc/lpc.go

PHONY: build install