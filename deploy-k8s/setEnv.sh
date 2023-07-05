# Description: Set environment variables for the k8sFileChurner application
export NUM_NAMESPACES=1 # int - Number of namespaces to create for testing
export NAMESPACE_PREFIX="test" # string - test namespace prefix
export STORAGE_CLASS_NAME="default" # string - ***only RWO storage 
export NUM_PVC_PER_NS=1 # string - number of PVC per Namespace 
export IMAGE_NAME="iuriikogan/k8sfilechurner:lastest" # string - default image or set to your own if you want to build the binary in from the makefile instead
export APP_SIZE_OF_PVC_GB=30 # int - size of PVCs (for each PVC a deployment will be created and a pod mounted
export APP_SIZE_OF_FILES_MB=100 # int - the size of files to be created per PVC
export APP_CHURN_PERCENTAGE=0.2 # float64 (0.6 = 60%) percentage of files to churn
export APP_CHURN_INTERVAL_MINUTES=60 # int - interval minutes to churn files