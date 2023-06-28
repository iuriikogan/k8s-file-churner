namespaces=$(kubectl get namespaces | grep test | awk '{print $1}')
for ns in $namespaces; do
  echo "Deleting namespace $ns"
  kubectl delete namespace $ns
done