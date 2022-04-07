VERSION=`git describe --tags`

LDFLAGS=-ldflags "-X github.com/vivekmurali/spidey/cmd.Version=${VERSION}"

app:
	go run .

build:
	go build ${LDFLAGS} .


clean:
	rm spidey.exe

test:
	go test -v ./...