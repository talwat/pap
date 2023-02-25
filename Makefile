PREFIX = $(HOME)/.local

build:
	mkdir -vp build
	go build -o build

install: build/pap
	install -Dm755 build/pap $(PREFIX)/bin/pap

clean:
	rm -rvf build