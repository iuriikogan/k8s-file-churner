# !#/bin/bash
export NUM_NAMESPACES=3 # int Number of namespaces to create for testing
export NAMESPACE_PREFIX="test" # test namespace prefix
# Set the image name for the busybox containe
export IMAGE_NAME="iuriikogan/dev:createtestenv" # string default image or set to your own if you want to build the binary in from the makefile instead
export NUM_PVC_PER_NS=3 # string number of PVC per Namespace 
# Set the PVC size, number of PVC per Namespace  (for each PVC a deployment will be created and a pod mounted, then files will be created and churned) the size of files 
export PVC_SIZE_GB=30 # int/float64 size of PVCs (for each PVC a deployment will be created and a pod mounted
export SIZE_OF_FILES_GB=1 # int/float64  the size of files to be created per PVC
export STORAGE_CLASS="managed-premium" # string only RWO storage classes TODO add logic for RWX storage classes

# Set the Churn params for the image
export CHURN_PERCENTAGE=20 # int percentage of files to churn