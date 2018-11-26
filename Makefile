#Go parameters
GO=go
GORUN=${GO} run
GOTEST=${GO} test -v

run:
	${GORUN}

test:
	${GOTEST} ./src/ipc