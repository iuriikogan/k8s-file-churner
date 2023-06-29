<<<<<<< HEAD
# bin/createTestEnv
=======
# to deploy to k8s 

  vi setEnv.sh ## set the required env variables to tune your test environment
  ./deploy-k8s/deploy.sh

  to cleanup the test environment

  ./deploy-k8s/destroy.sh

# ./createTestEnv
>>>>>>> 541b92d (merge)
Go bin which writes a number of files with random data to /data/ directory

# Create bin from Makefile 

## make build 
  build binary and copies to /bin/${BINARY_NAME}

## make run
  make build and run the binary

## make clean
  delete the binary in /bin/${BINARY_NAME}


<<<<<<< HEAD
# Run container locally
  podman build -t <repository:tag> .
  podman run <repository:tag>
  
# Deploy-k8s
  ./deploy-k8s/deploy.sh
  deploys x deployments with x replicas/pvcs in x namespaces with x storage class 

  ./deploy-k8s/detroy.sh 
  deletes all the namespaces beginning with test


## clean up local test file
  
  ./utils/deleteTestFiles.sh
  cleans up test files from data/ dir from running the image locally
=======
  ./deleteTestFiles.sh
>>>>>>> 541b92d (merge)
