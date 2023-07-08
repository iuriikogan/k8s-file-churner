export NUM_NAMESPACES=1
export NAMESPACE_PREFIX="new-test"
export STORAGE_CLASS_NAME="managed-premium" # only RWO Storage classes
export NUM_PVC_PER_NS=3
export IMAGE_NAME="iuriikogan/k8sfilechurner:v1"
export APP_SIZE_OF_PVC_GB=30
export APP_SIZE_OF_FILES_MB=99
export APP_CHURN_PERCENTAGE=0.2 # this value should be a float64 representing % i.e 0.2 = 20%
export APP_CHURN_INTERVAL_MINUTES="15m" # this should be a duration in min as string i.e. "*m"
