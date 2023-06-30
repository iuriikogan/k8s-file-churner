# ./createTestEnv
    Go bin which writes a number of files with random data to /data/ directory


# to deploy to k8s 

    git clone https://github.com/iuriikogan/createTestEnv-image.git

    cd deploy-k8s

    vi setEnv.sh ## set the required env variables to tune your test environment

    ./deploy.sh
    deploys x deployments with x replicas/pvcs in x namespaces with x storage class 

    ./detroy.sh 
  
    deletes all the namespaces beginning with $NAMESPACE_PREFIX

# Create bin from Makefile 

## make build 
    build binary and copies to /bin/${BINARY_NAME}

## make run
    make build and run the binary

## make clean
    delete the binary in /bin/${BINARY_NAME} and any local test files in /data dir as well as delete all the namespace with $NAMESPACE_PREFIX

# Run container locally
    podman build -t <repository:tag> .
    podman run <repository:tag>

## clean up local test file
  
    ./utils/deleteTestFiles.sh
    cleans up test files from data/ dir from running the image locally
