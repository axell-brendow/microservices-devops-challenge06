# microservices-devops-challenge06

Course "[Architecting and Developing modern and scalable applications based on Microservices (pt-br)](https://drive.google.com/file/d/1JXXmYhfi-Sk0zwiEbBIEswDC6AFeyBlD/view?usp=sharing)".

- Http server in Go with CI and tests in GCP
- Apply a Horizontal Pod Autoscaler in a Go application under high usage
    > Docker Image: [axell13/go-hpa](https://hub.docker.com/r/axell13/go-hpa)

To create this Deployment in your cluster:

```
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
kubectl apply -f hpa.yaml
```

To test the auto scaler, run `kubectl run -it loader --image=busybox /bin/sh`

And then, inside busybox, run `while true; do wget -q -O- http://go-hpa.default.svc.cluster.local; done`

In another terminal, you can see auto scaler working with `kubectl get hpa`. If you wanna watch the metrics, run `watch kubectl get hpa`

## Other challenges

- [Nginx + PHP-FPM + Redis + MySQL with docker-compose](https://github.com/axell-brendow/microservices-devops-challenge01)
- [Create Hello World in Go with a Docker image less than 2MB](https://github.com/axell-brendow/microservices-devops-challenge02)
- [Apply CI (Continuous Integration) to test and create a docker image of a simple Go application](https://github.com/axell-brendow/microservices-devops-challenge04)
- [3 Nginx replicas with a Load Balancer; MySQL with Secret Object and PersistentVolume; Http server in Go with CI and tests in GCP](https://github.com/axell-brendow/microservices-devops-challenge05)
- [Apply a Horizontal Pod Autoscaler in a Go application under high usage](https://github.com/axell-brendow/microservices-devops-challenge06)
