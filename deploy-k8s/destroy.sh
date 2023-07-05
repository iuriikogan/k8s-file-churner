#!/bin/bash
namespaces=$(kubectl get namespaces | grep $NAMESPACE_PREFIX | awk '{print $1}')
for ns in $namespaces; do
  echo "Deleting namespace $ns"
  kubectl delete namespace $ns
done