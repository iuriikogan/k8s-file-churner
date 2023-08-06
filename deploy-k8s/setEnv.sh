export NUM_NAMESPACES=1
export NAMESPACE_PREFIX="test-v1"
export STORAGE_CLASS_NAME="do-block-storage" # only RWO Storage classes
export NUM_PVC_PER_NS=1
export IMAGE_NAME="iuriikogan/k8sfilechurner:latest"
export APP_SIZE_OF_PVC_GB=10
export APP_SIZE_OF_FILES_MB=100
export APP_CHURN_PERCENTAGE=0.2 # this value should be a float64 representing % i.e 0.2 = 20%
export APP_CHURN_INTERVAL_MINUTES="5m" # this should be a duration in min as string i.e. "*m"
