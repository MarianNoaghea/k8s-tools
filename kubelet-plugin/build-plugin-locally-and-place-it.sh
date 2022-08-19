#!/bin/bash

cd ~/kubernetes/go-plugins/demo1-plugins;
make;
sudo cp plugin1.so /usr/bin/plugin1.so;
echo 'The library was built right now see below in /usr/bin'
ls -lrt /usr/bin/ | grep plugin1.so;
cd ~/k8s-tools/kubelet-plugin;

