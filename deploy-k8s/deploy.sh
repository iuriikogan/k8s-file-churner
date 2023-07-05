#!/bin/bash
echo "-------Deploying K8sFileChurner"
starttime=$(date +%s)
. ./setenv.sh
echo 'starting to create deployments'
# Loop to create namespaces and deployments
for ((i=1; i<=NUM_NAMESPACES; i++))
do
  NAMESPACE="$NAMESPACE_PREFIX-$i"
  # Create the namespace
  kubectl create namespace $NAMESPACE
  # Create the configmap shared by all the deployments as mounted volume
  kubectl -n $NAMESPACE apply -f - <<EOF
    apiVersion: v1
    kind: ConfigMap
    metadata:
      name: app-cm
    data:

EOF    
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
          storage: ${APP_PVC_SIZE_GB}
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
                mountPath: app/
              resources:
                requests:
                  memory: 1Gi
                  cpu: 0.5
                limits:  
                  memory: 1Gi
                  cpu: 1
              command: ["sh"]                
          volumes:
          - name: data
            persistentVolumeClaim:
              claimName: pvc-${NAMESPACE}-${j}
          - name: config
            configMap:
              name: app-cm
              items:
              - key: APP_SizeOfFileMB
                path: app.SizeOfFileMB
EOF
done

done
endtime=$(date +%s)
duration=$(( $endtime - $starttime ))
echo "-------Finished deploying K8sFileChurner in $(($duration / 60)) minutes $(($duration % 60)) seconds."
