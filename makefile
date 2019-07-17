all: build

build:
        go build -o adoc .
install:
        cp adoc /usr/local/bin
clean:
        rm adoc /usr/local/bin/adoc
