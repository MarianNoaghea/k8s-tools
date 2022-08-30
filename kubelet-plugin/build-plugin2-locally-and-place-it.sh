#!/bin/bash

cd ~/kubernetes;
make plugin WHAT=cmd/kubelet/demo2-plugins --env CGO_ENABLED=1 --env GO111MODULES=on GOFLAGS=-v
sudo cp plugin2.so /usr/bin
mv plugin2.so myplugins
echo 'Plugin was built right now, see below in /usr/bin';
ls -lrt /usr/bin/ | grep plugin2.so;
cd ~/k8s-tools/kubelet-plugin
