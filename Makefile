install:
	go get github.com/dlarocque/habi

precommit:
	pre-commit run --all-files

fmt:
	go fmt ./...
