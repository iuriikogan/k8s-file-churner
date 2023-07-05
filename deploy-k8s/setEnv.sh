#!/bin/bash
export NUM_NAMESPACES=1
export NAMESPACE_PREFIX="test"
export STORAGE_CLASS_NAME="managed-premium"
export NUM_PVC_PER_NS=3
export IMAGE_NAME="iuriikogan/k8sfilechurner:lastest"
export APP_SIZE_OF_PVC_GB=40
export APP_SIZE_OF_FILES_MB=50
export APP_CHURN_PERCENTAGE=0.2 # this value should be a float64 representing % i.e 0.2 = 20%
export APP_CHURN_INTERVAL_MINUTES="30s" # this should be a duration i.e. 5m, 1h, etc