ip=`kubectl get ingress -o wide | grep -v NAME | awk {'print $4'}`
echo "minikube ingress ip" $ip

curl -X POST $ip/entity/ -H "Content-Type: application/json" -d '{"id": 1, "value": "test123"}'
curl -X GET $ip/entity/1
