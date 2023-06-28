#!/bin/bash

# Set the number of namespaces to create

NUM_NAMESPACES=3
NAMESPACE_PREFIX="test"

# Set the image name for the busybox container

IMAGE_NAME="iuriikogan/dev:createtestenv"

# Set the PVC size number of PVC (replicas of deployment)and storage class

PVC_SIZE="30Gi"
NUMBER_OF_FILES=30
SIZE_OF_FILES="1GB"
NUM_PVC_PER_NS=3
STORAGE_CLASS="managed-premium"

# Loop to create namespaces and deployments
for ((i=1; i<=NUM_NAMESPACES; i++))
do
  NAMESPACE="$NAMESPACE_PREFIX-$i"
  echo "*************************************************"
  echo "tried creating namespace: $NAMESPACE"
  echo "*************************************************"
  # Create the namespace
  kubectl create namespace $NAMESPACE
  # Create the PVCs and deployments for $NUM_OF_PVC_PER_NS
  for ((j=1; j<=$NUM_PVC_PER_NS; j++)); do
    # Create the PVC for each Deploy 
    kubectl -n $NAMESPACE apply -f - <<EOF
    apiVersion: v1
    kind: PersistentVolumeClaim
    metadata:
      name: pvc-${NAMESPACE}-${j}
    spec:
      accessModes:
        - ReadWriteOnce
      resources:
        requests:
          storage: ${PVC_SIZE}
      storageClassName: ${STORAGE_CLASS}
EOF

# Create the deployment
    kubectl -n $NAMESPACE apply -f - <<EOF
    apiVersion: apps/v1
    kind: Deployment
    metadata:
      name: test-deploy-${NAMESPACE}-${j}
    spec:
      replicas: 1
      selector:
        matchLabels:
          app: test
      template:
        metadata:
          labels:
            app: test
        spec:
          containers:
            - name: test-pod-${NAMESPACE}-${j}
              image: ${IMAGE_NAME}
              imagePullPolicy: Always
              volumeMounts:
              - name: data
                mountPath: app/data/
              resources:
                requests:
                  memory: 1Gi
                  cpu: 0.5
                limits:  
                  memory: 1Gi
                  cpu: 1
          volumes:
          - name: data
            persistentVolumeClaim:
              claimName: pvc-${NAMESPACE}-${j}
EOF
done
  echo "*************************************************"
  echo "Tried to create pvcs and deployments in $NAMESPACE"
  echo "*************************************************"
done
