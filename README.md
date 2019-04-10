# lfs - cp
latitude financial services - cloud platform coding test

Instructions:
1. Run the app directly as follows:
```bash
docker pull sutoor/lfs.cp.test:1.0
docker run --rm sutoor/lfs.cp.test:1.0 lfs maxprofit --prices "90,100,2,25,3,4,27,26"
#run unit tests as follows:
docker run --rm -w /go/src/lfs/stock sutoor/lfs.cp.test:1.0 go test
```

2. It can also be built as a docker container & then run as follows: 
```bash
docker build -t lfs.cp.test:1.0 .
docker run --rm lfs.cp.test:1.0 lfs maxprofit --prices "90,100,2,25,3,4,27,26"
#and run unit tests as follows:
docker run --rm -w /go/src/lfs/stock lfs.cp.test:1.0 go test
```

3. The app is written in Go. It can be built from source as a native app by `cd`ing into top level directory (containing this readme) and:
```bash
#make sure the repo is checked out under $HOME/go/src and then execute:
go get ./...
go install
go test
#lfs command should now be available(provided $GOPATH/bin is in your $PATH env variable)
lfs -h
lfs maxprofit -h
lfs maxprofit -prices "90,100,2,25,3,4,27,26"
```

4. There's also a kubernetes deployment file if it needs to be deployed on a kubernetes cluster. That can be done as follows:
```bash
kubectl apply -f k8sdeploy.yaml
#to scale for more prod load
kubectl get deploy
kubectl scale deploy lfs-profit --replicas=5
```
