PREFIX = $(HOME)/.local

build:
	mkdir -vp build
	go build -o build

install: build/pap
	install -m755 build/pap $(PREFIX)/bin/pap

clean:
	rm -rvf build