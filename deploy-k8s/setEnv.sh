# !#/bin/bash
export NUM_NAMESPACES=3 # int - Number of namespaces to create for testing
export NAMESPACE_PREFIX="test" # string - test namespace prefix
# Set the image name for the busybox containe
export IMAGE_NAME="iuriikogan/dev:createtestenv" # string - default image or set to your own if you want to build the binary in from the makefile instead
export NUM_PVC_PER_NS=3 # string - number of PVC per Namespace 
# Set the PVC size, number of PVC per Namespace  (for each PVC a deployment will be created and a pod mounted, then files will be created and churned) the size of files 
export APP_PVC_SIZE_GB=30 # int - size of PVCs (for each PVC a deployment will be created and a pod mounted
export APP_SIZE_OF_FILES_MB=100 # int - the size of files to be created per PVC
export STORAGE_CLASS_NAME="managed-premium" # string - ***only RWO storage 
export APP_CHURN_PERCENTAGE=0.2 # float64 (0.6 = 60%) percentage of files to churn
export APP_CHURN_INTERVAL_MINUTES=60 # int - interval minutes to churn files