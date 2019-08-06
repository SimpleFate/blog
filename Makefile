TARGET=blog
PROJECT=blog

ARCH=amd64

buildLocal:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ${TARGET} "."

buildLinux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${TARGET} "."

uploadAll:
	ssh corey@myserver "cd project/blog/; rm -r static; rm -r view; rm ${TARGET}"
	scp  ${TARGET} corey@myserver:/home/corey/project/blog/
	scp -r static corey@myserver:/home/corey/project/blog/
	scp -r view corey@myserver:/home/corey/project/blog/

clean:
	rm -rf ${TARGET}

.PHONY: clean buildLocal buildLinux uploadAll