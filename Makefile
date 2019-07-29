TARGET=blog

OS=linux
#OS=windows
#OS=darwin

ARCH=amd64

build:
	CGO_ENABLED=0 GOOS=${OS} GOARCH=amd64 go build -o bin/${OS}/${TARGET} "."

upload:
	scp bin/${OS}/${TARGET} corey@myserver:/home/corey/blog

run:
	ssh corey@myserver "cd ~/blog; ./blog"

clean:
	rm -rf bin

.PHONY: clean build