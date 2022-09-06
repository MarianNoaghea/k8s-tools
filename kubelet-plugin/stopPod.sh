#!/bin/sh
kubectl delete -f $1 --force --grace-period=0
