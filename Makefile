export GO111MODULE=on

dep:
	go mod download

test:
	go test -v
