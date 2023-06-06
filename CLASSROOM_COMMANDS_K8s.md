# Commands for Classroom Session

> This is where I store my commands for the classroom session for Intro to Kubernetes. Just in case someone wants to follow along or they want to try it later. These commands should work as long as you are in a bash or similar terminal (git-bash or WSL for Windows, term or iterm2 for Mac, or any basic terminal on any Linux OS with kubectl installed).

## Prereqs

* Rancher Desktop installed on your laptop desktop computer.

> I am using Rancher Desktop on my Mac to run these commands, but Docker Desktop with Kubernetes enabled works fine too.

## k8s template

```
kubectl 
```

## k8s pods

```
cd intro-2-pods

kubectl apply -f py-samp-pod.yaml

kubectl get pods

kubectl logs py-samp

kubectl exec -it py-samp -- /bin/bash

exit

kubectl port-forward pod/py-samp 8082:8082

kubectl delete -f py-samp-pod.yaml

kubectl get pods
```

## k8s deployments

```
cd intro-2-deploy

kubectl apply -f py-samp-deploy.yaml

kubectl get deploy

kubectl get pods

kubectl logs <py-samp-unique-id>

kubectl port-forward deploy/py-samp 8082:8082

kubectl delete -f py-samp-deploy.yaml

kubectl get deploy

kubectl get pods
```

## k8s services

```
kubectl apply -f intro-2-svc/

kubectl get svc

kubectl get deploy

kubectl delete -f intro-2-svc/
```

## k8s namespaces

```
kubectl apply -f intro-2-namespaces/

kubectl get deploy

kubectl get svc

kubectl delete -f intro-2-namespaces/
```

## k8s ingress

```
kubectl apply -f intro-2-ingress/

kubectl get deploy

kubectl get svc

kubectl get ing

kubectl delete -f intro-2-ingress/
```
