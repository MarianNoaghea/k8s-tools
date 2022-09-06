#!/bin/bash

cd ~/kubernetes;
make plugin WHAT=cmd/kubelet/demo1-plugins --env CGO_ENABLED=1 --env GO111MODULES=on GOFLAGS=-v
sudo mv plugin1.so /etc/kubernetes/cci
echo 'Plugin was built right now, see below in /etc/kubernetes/cci';
ls -lrt /etc/kubernetes/cci | grep plugin1.so;
cd ~/k8s-tools/kubelet-plugin
