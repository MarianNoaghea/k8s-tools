#!/bin/sh
kubectl apply -f pod-cpu.yaml
sleep 4
taskset -c -p $(pgrep stress)
