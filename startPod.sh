#!/bin/sh
kubectl apply -f pod-cpu.yaml
sleep 1
taskset -c -p $(pgrep stress)
