# k8s-file-churner

The **k8s-file-churner** is a tool designed to simulate file churn in a Kubernetes environment. It creates and churns a specified number of files within a directory, allowing you to test the behavior of your backup and storage systems or applications that interact with the files.

## Features

- Creates a specified number of files with random data.
- Churns a percentage of the files by deleting them and recreating them with new random data.
- Supports customization of file size, churn percentage, and churn interval through a configuration file.
- Concurrent file creation and churn operations using goroutines.
- Outputs statistics such as the size of each file, the number of files created, and the time taken for the operation.

## Prerequisites

- Go 1.16 or higher installed on your machine.
- Kubernetes cluster configured and accessible via `kubectl`.

## Limitations

- Currently only RWO storage classes are supported

## Getting Started

1. Clone the repository:

   ```shell
   git clone https://github.com/iuriikogan/k8s-file-churner.git
   ```
2. Navigate to the created directory

   ```shell
   cd deploy-k8s/
   ```

3. Set the env variables in setenv.sh

   ```shell
   vi setenv.sh
   ```
4. deploy to k8s (**Double check you are in the right context**)

   ```shell
   ./deploy-k8s
   ```
5. delete the test namespaces
 
    ```shell
    ./destroy.sh
    ```
## Build from the source

1. Navigate to the root project directory and build the binaries from the Makefile

2. This will generate the binary file 'bin/k8sFileChurner'
   ```shell
   make build
   ```

3. This will execute the application binary
   ```shell
   make run
   ```

4. This will run the tests for the project.
   ```shell
   make test
   ```
5. This will remove the generated binary, delete test files, and clean up any Kubernetes resources created during testing.
   ```shell
   make clean
   ```
## Contact Me!
I'm always interested in collaborating and I'd love to hear your feedback!! - koganiurii@gmail.com



