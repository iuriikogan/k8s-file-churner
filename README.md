# k8s-file-churner

The **k8s-file-churner** is a tool designed to simulate file churn in a Kubernetes environment. It creates and churns a specified number of files within a directory, allowing you to test the behavior of your storage systems or applications that interact with the files.

## Features

- Creates a specified number of files with random data.
- Churns a percentage of the files by deleting them and recreating them with new random data.
- Supports customization of file size, churn percentage, and churn interval through a configuration file.
- Concurrent file creation and churn operations using goroutines.
- Outputs statistics such as the size of each file, the number of files created, and the time taken for the operation.

## Prerequisites

- Go 1.16 or higher installed on your machine.
- Kubernetes cluster configured and accessible via `kubectl`.

## Getting Started

1. Clone the repository:

   ```shell
   git clone https://github.com/iuriikogan/k8s-file-churner.git
   cd deploy-k8s/
2. Set the env variables in setenv.sh

   ```shell
   vi setenv.sh

3. deploy to k8s (**Double check you are in the right context**)

   ```shell
   ./deploy-k8s

4. delete the test namespaces
 
    ```shell
    ./destroy.sh