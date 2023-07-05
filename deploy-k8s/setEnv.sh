# Description: Set environment variables for the k8sFileChurner application
export NUM_NAMESPACES=1
export NAMESPACE_PREFIX="test"
export STORAGE_CLASS_NAME="managed-premium"
export NUM_PVC_PER_NS=1
export IMAGE_NAME="iuriikogan/k8sfilechurner:lastest"
export SIZE_OF_PVC_GB=50
export SIZE_OF_FILES_MB=100
export CHURN_PERCENTAGE=0.2
export CHURN_INTERVAL_MINUTES=60
