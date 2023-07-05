#!/bin/bash
export NUM_NAMESPACES=3 
export NAMESPACE_PREFIX="test"
export STORAGE_CLASS_NAME="managed-premium" # only RWO Storage classes
export NUM_PVC_PER_NS=3
export IMAGE_NAME="iuriikogan/k8sfilechurner:lastest"
export APP_SIZE_OF_PVC_GB=30
export APP_SIZE_OF_FILES_MB=999
export APP_CHURN_PERCENTAGE=0.2 # this value should be a float64 representing % i.e 0.2 = 20%
export APP_CHURN_INTERVAL_MINUTES="1m" # this should be a duration in min as string i.e. "*m"