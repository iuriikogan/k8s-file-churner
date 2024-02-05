BINARY_NAME=k8sFileChurner

build: 

	go build -o ./bin/$(BINARY_NAME) ./main.go 

run: build
	
	./bin/$(BINARY_NAME)

test: 

	go test ./... -v --count=1

bench:

	go test ./... -bench=. -benchmem -v

clean:

	rm -f ./bin/$(BINARY_NAME)
	./deleteTestFiles.sh
	./deploy-k8s/destroy.sh

lint: 

	golint ./...
