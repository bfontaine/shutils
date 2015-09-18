# Shutils Makefile

PREFIX ?= .
BIN    ?= $(PREFIX)/bin

EXES = $(BIN)/sum

install: $(EXES)

$(BIN):
	mkdir -p $@

$(BIN)/%: $(BIN)
	go build -o $@ $(notdir $@).go
