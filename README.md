# HOW IT WORKS?
## Clone Repo
```
$ git clone git@github.com:life1347/googlepusbsub-exmaple.git $GOPATH/src/tachingchen.com/googlePubSub
```

## Restore Dependencies
```
$ godep restore
```

## Pull Model
```
$ GOOGLE_CLOUD_PROJECT="<google project name>" go run pull.go
```

## Push Model
### Build docker image
* Place domain SSL `certificate.crt` and `private.key` in asset directory
* Place `crt.json` in asset directory with json file download from API Console Credentials page. (Please refer [Google Dev Guide](https://developers.google.com/identity/protocols/application-default-credentials))
* (Optional) Replace `ca-certificates.crt` with the ca-certificates.crt, which can be found under the /etc/ssl/certs/ in any linux distribution, in asset directory.

```
$ cd push
# build.sh will upload built image to google image registry automatically
$ GOOGLE_CLOUD_PROJECT="<google project name>" ./build.sh
```

### K8S
* Replace configuration in k8s/pubsub-deployment.yaml
* Point your domin name (endpoint) to `EXTERNAL-IP`

```
$ kubectl apply -f k8s/pubsub-deployment.yaml
$ kubectl apply -f k8s/pubsub-svc.yaml
$ kubectl get svc
NAME                      CLUSTER-IP     EXTERNAL-IP       PORT(S)            AGE
pubsub-service            10.3.246.101   x.x.x.x           443/TCP            39s
```

## Publisher
```
# Once the push/pull client is ready
$ GOOGLE_CLOUD_PROJECT="<google project name>" go run publisher.go
```
