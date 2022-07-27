#!/bin/sh
kubectl run chosenone --image=nginx -o yaml --dry-run=client > mypod.yaml
kubectl apply -f mypod.yaml
taskset -c -p $(ps -ef | pgrep nginx |sed -n '2 p')
kubectl delete -f mypod.yaml --force --grace-period=0
