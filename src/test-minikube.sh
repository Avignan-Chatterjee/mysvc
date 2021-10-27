url=`minikube service mysvc-service | grep http | awk {'print $8'}`
echo "minikube service url" $url

curl -X POST $url/entity -H "Content-Type: application/json" -d '{"id": 1, "value": "test123"}'
curl -X GET $url/entity/1



