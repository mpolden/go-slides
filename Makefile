all: deps

deps:
	go get golang.org/x/tools/cmd/present

run:
	bin/present
