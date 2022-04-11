app:
	go run .

b:
	go build .


clean:
	rm spidey.exe

test:
	go test -v ./...