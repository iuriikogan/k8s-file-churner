# bin/createTestEnv
Go bin which writes a number of files with random data to /data/ directory

# Create from Makefile 
## make build 
  build binary and copies to /bin/${BINARY_NAME}

## make run
  make build and run the binary

## make clean
  delete the binary in /bin/${BINARY_NAME}

## clean up local test file
  to cleanup test files locally, run the following command

./utils/deleteTestFiles.sh
 
 to be removed after deleteFiles.go is written
