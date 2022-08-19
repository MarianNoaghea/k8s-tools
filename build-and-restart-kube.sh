#!/bin/bash

~/kubernetes/build/run.sh make DBG=1 --env  CGO_ENABLED=1  kubelet
sudo rm /usr/bin/kubelet
sudo cp ~/kubernetes/_output/dockerized/bin/linux/amd64/kubelet /usr/bin/
systemctl restart kubelet
systemctl status kubelet
