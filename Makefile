.PHONY:	clean fmt test vet

fmt:
	go fmt .

test:
	go test .

dep:
	go get github.com/gorilla/mux

clean:
	go clean
