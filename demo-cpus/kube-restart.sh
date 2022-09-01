#!/bin/sh
systemctl daemon-reload
sudo rm /var/lib/kubelet/cpu_manager_state
systemctl restart kubelet.service
echo 'wait some time then $ systemctl status kubelet.service'

