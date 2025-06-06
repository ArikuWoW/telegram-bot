CURDIR=$(shell pwd)
BINDIR=${CURDIR}/bin
GOVER := $(shell go version | perl -nle '/go(\d+\.\d+)/; print $$1;')
MOCKGEN := $(BINDIR)/mockgen_$(GOVER)
SMARTIMPORTS := $(BINDIR)/smartimports_$(GOVER)
LINTBIN := $(BINDIR)/lint_$(GOVER)_$(LINTERVER)
PACKAGE := github.com/ArikuWoW/telegram-bot/cmd/bot

all: format build test lint

build: bindir
	go build -o $(BINDIR)/bot $(PACKAGE)

test:
	go test ./...

run:
	go run $(PACKAGE)

generate: install-mockgen
	$(MOCKGEN) -source=internal/model/messages/incoming_msg.go -destination=internal/mocks/messages/messages_mocks.go

lint: install-lint
	${LINTBIN} run

precommit: format build test lint
	echo "OK"

bindir:
	mkdir -p ${BINDIR}

format: install-smartimports
	${SMARTIMPORTS} -exclude internal/mocks

install-mockgen: bindir
	test -f ${MOCKGEN} || \
		(GOBIN=${BINDIR} go install github.com/golang/mock/mockgen@v1.6.0 && \
		mv ${BINDIR}/mockgen ${MOCKGEN})

install-lint: bindir
	test -f ${LINTBIN} || \
		(GOBIN=${BINDIR} go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.0.2 && \
		mv ${BINDIR}/golangci-lint ${LINTBIN})

install-smartimports: bindir
	test -f ${SMARTIMPORTS} || \
		(GOBIN=${BINDIR} go install github.com/pav5000/smartimports/cmd/smartimports@latest && \
		mv ${BINDIR}/smartimports ${SMARTIMPORTS})

docker-run:
	sudo docker compose up