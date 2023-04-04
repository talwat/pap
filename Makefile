PREFIX = $(HOME)/.local

build:
	mkdir -vp build
	go build -o build

install: build/pap*
	mkdir -p $(PREFIX)/bin
	mv build/pap* $(PREFIX)/bin
	chmod +x $(PREFIX)/bin/pap*

clean:
	rm -rvf build