#Go parameters
GO=go
GORUN=${GO} run
GOTEST=${GO} test -v
GOBUILD=${GO} build

run:
	${GORUN} src/cgss/cgss.go

test:
	${GOTEST} ./src/ipc

append:
	${GOTEST} ./src/splice-append