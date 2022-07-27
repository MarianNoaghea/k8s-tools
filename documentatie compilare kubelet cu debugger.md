# Compile kubelet with debugger in linux terminal

Steps
-


From director /home/ixia:

1. cd kubernetes/
2. ./build/run.sh make DBG=1 kubelet (**if an error occurs, edit file /etc/resolv.conf -> add the line nameserver 8.8.8.8 before the line nameserver 127.0.0.53**)
3. sudo cp _output/dockerized/bin/linux/amd64/kubelet /usr/bin/ (**copy the file in /usr/bin -> to figure out in which director you should copy the file, press command *which kubelet***)
4. systemctl restart kubelet
5. systemctl status kubelet (**after restaring kubelet, you should check the status -> it should be active**)
6. ps aux | grep kubelet **(this command will give us the pid for the process which runs kubelet) -> we will use it at step 8**
7. **install dlv -> switch to sudo user, then type the command:** go install github.com/go-delve/delve/cmd/dlv@latest
8. dlv attach $PID /home/ixia/kubernetes/_output/dockerized/bin/linux/amd64/kubelet --headless --listen=0.0.0.0:2345 --log --api-version 2 **(here, you should replace $PID with the pid found at step 6)**
9. dlv connect localhost:2345
