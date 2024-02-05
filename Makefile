MAIN_GO=zb.go
BIN_NAME=zb

.PHONY: help test
help:
	@echo "usage: make <option>"
	@echo "options and effects:"
	@echo "    help                     : Show help"
	@echo "    test                     : Test ..."

test:
	@echo "test ..."

.PHONY: push build mac linux image
push:
	@git add .
	git commit -m "自动push"
	git push origin main

build:
	go build -ldflags="-s -w" ${MAIN_GO}
	$(if $(shell command -v upx), upx ${BIN_NAME})

mac:
	GOOS=darwin go build -ldflags="-s -w" -o ${BIN_NAME}-darwin ${MAIN_GO}
	$(if $(shell command -v upx), upx ${BIN_NAME}-darwin)

win:
	GOOS=windows go build -ldflags="-s -w" -o ${BIN_NAME}.exe ${MAIN_GO}
	$(if $(shell command -v upx), upx ${BIN_NAME}.exe)

linux:
	GOOS=linux go build -ldflags="-s -w" -o ${BIN_NAME}-linux ${MAIN_GO}
	$(if $(shell command -v upx), upx ${BIN_NAME}-linux)