#!/bin/bash
echo "-------Deploying K8sFileChurner"
starttime=$(date +%s)
source ./setenv.sh
echo "Starting to create deployments this shouldn't take long..."

# Loop to create namespaces and deployments
for ((i=1; i<=$NUM_NAMESPACES; i++)); do
  NAMESPACE="$NAMESPACE_PREFIX-$i"

  # Create the namespace
  kubectl create namespace $NAMESPACE

  # Create the ConfigMap
  cat <<EOF | kubectl apply -n $NAMESPACE -f -
apiVersion: v1
kind: ConfigMap
metadata:
  name: config
data:  
    APP_SIZE_OF_FILES_MB: "${APP_SIZE_OF_FILES_MB}"
    APP_SIZE_OF_PVC_GB: "${APP_SIZE_OF_PVC_GB}"
    APP_CHURN_PERCENTAGE: "${APP_CHURN_PERCENTAGE}"
    APP_CHURN_INTERVAL_MINUTES: "${APP_CHURN_INTERVAL_MINUTES}"
EOF

  # Create the PVCs and deployments for $NUM_PVC_PER_NS
  for ((j=1; j<=$NUM_PVC_PER_NS; j++)); do
    # Create the PVC for each Deployment
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
      storage: ${SIZE_OF_PVC_GB}Gi
  storageClassName: ${STORAGE_CLASS}
EOF

    # Create the Deployment
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
              mountPath: app/
          resources:
            requests:
              memory: 1Gi
              cpu: 0.5
            limits:
              memory: 1Gi
              cpu: 1
          envFrom:
          - configMapRef:
              name: config
      volumes:
        - name: data
          persistentVolumeClaim:
            claimName: pvc-${NAMESPACE}-${j}
EOF

  done
done

# Wait for pods to start running
echo "Waiting for pods to start running..."
for ((i=1; i<=$NUM_NAMESPACES; i++)); do
  NAMESPACE="$NAMESPACE_PREFIX-$i"

  # Wait for pods in the namespace to start running
  kubectl -n $NAMESPACE wait --for=condition=ready pod --all --timeout=300s
  endtime=$(date +%s)
  duration=$(( $endtime - $starttime ))
  echo "-------Finished deploying K8sFileChurner in $(($duration / 60)) minutes $(($duration % 60)) seconds."
done