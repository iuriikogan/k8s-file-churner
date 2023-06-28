BINARY_NAME=createTestEnv

build: 

	go build -o ./bin/$(BINARY_NAME) ./main.go

run: build
	
	./bin/$(BINARY_NAME)

clean:

	rm -f ./bin/$(BINARY_NAME)
	./utils/deleteTestFiles.sh

	