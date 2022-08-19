#!/bin/bash

cd ~/kubernetes;
make all WHAT=cmd/kubelet GOFLAGS=-v
sudo cp ~/kubernetes/_output/bin/kubelet /usr/bin/
echo 'Kubelet was built right now, see below in /usr/bin';
ls -lrt /usr/bin/ | grep kubelet;
cd ~/k8s-tools/kubelet-plugin
