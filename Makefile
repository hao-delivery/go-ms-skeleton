SERVICE=app

.PHONY: local

local:
	go build -o $(SERVICE) main.go
	clear		
	./$(SERVICE) run


.PHONY: run

run:		
	./$(SERVICE) run


.PHONY: build

build:
	go build -o $(SERVICE) main.go
	


count:
	find . -name tests -prune -o -type f -name '*.go' | xargs wc -l

.DEFAULT_GOAL := local