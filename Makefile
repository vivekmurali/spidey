VERSION=`git describe --tags`

LDFLAGS=-ldflags "-X github.com/vivekmurali/spidey/cmd.Version=${VERSION}"


build:
	go build ${LDFLAGS} .


clean:
	rm spidey.exe
