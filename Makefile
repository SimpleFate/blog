TARGET=blog
PROJECT=blog

ARCH=amd64

buildLocal:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ${TARGET} "."

buildLinux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${TARGET} "."

uploadAll:
	scp -r ../${PROJECT} corey@myserver:/home/corey/project/

clean:
	rm -rf ${TARGET}

.PHONY: clean buildLocal buildLinux uploadAll