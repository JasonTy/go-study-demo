#Go parameters
GO=go
GORUN=${GO} run
GOTEST=${GO} test -v
GOBUILD=${GO} build

run:
	${GORUN} main.go

test:
	${GOTEST} ./src/ipc

build:
	${GOBUILD} main.go

append:
	${GOTEST} ./src/splice-append