# Shutils Makefile

PREFIX ?= /usr/local
BIN    ?= $(PREFIX)/bin

EXES = $(BIN)/sum

install: $(EXES)

uninstall:
	rm -f $(EXES)

$(BIN):
	mkdir -p $@

$(BIN)/%: $(BIN) %.go
	go build -o $@ $(notdir $@).go
