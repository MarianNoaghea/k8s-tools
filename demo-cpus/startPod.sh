#!/bin/sh
kubectl apply -f $1
sleep 4
taskset -c -p $(pgrep stress)
