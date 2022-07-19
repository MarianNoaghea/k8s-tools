# K8s build + dockerised documentation

Build
-

1. Open terminal -> go to kubernetes directory
2. \$ make clean
3. \$ ./build/run.sh make

Find the path to the binary file kube-apiserver:
-

1. My way is:
2. \$ tree . > out.txt
3. \$ gedit out.txt &
4. CTRL + F look for "kube-apiserver" \
 Now you know the path
5. the path is: /kubernetes/_output/dockerized/bin/linux/amd64/kube-apiserver

Make a work directory
-

1. \$ mkdir ~/work-dir
2. \$ cp _output/dockerized/bin/linux/amd64/kube-apiserver ~/work-dir/
3. \$ cd ~/work-dir

Docker
-

1. $ code Dockerfile
2. Paste
```docker
FROM k8s.gcr.io/kube-apiserver:v1.24.2
COPY /kube-apiserver /usr/local/bin/kube-apiserver

```
3. Build your image\
$ docker image build -t k8s.gcr.io/kube-apiserver:v1.24.2_yourName .
4. Archive your image\
$ docker save k8s.gcr.io/kube-apiserver:v1.24.2_yourName > archiveName.tar.gz
5. Use containerd CLI to import your image\
$ ctr -n=k8s.io images import archiveName.tar.gz 
6. Show your images\
$crictl images -a\
You can notice that your "kube-apiserver: v1.24.2_yourName" is pressent.

# Logging in Kubernetes

Hello, folks! In today's tutorial, I will show you how to add to kubelet the --log-file option.
Let's get started!
I will present the commands, along with the related explanations.


Steps
-

1. **Explanation** : Switch to sudo user \
   **Command**: *sudo su*

2.  **Explanation**: Open the file from the path: /etc/systemd/system/ kubelet.service.d/10-kubeadm.conf \
                 Use your favourite text editor. \
    **Command**: *vim /etc/systemd/system/kubelet.service.d/10-kubeadm.conf*

3.  **Explanation**: In the opened file, you will see two lines including the keyword "Environment". \
                 You cand add the string "--log-file=/tmp/kubelet.log -v=8" at the end of any of this two lines. \
                 I added it at the end of the second line. \
                 The log messages can be seen in the file from the path /tmp/kubelet.log. If you prefer a different location,
                 feel free to put the log messages where you want. \ The -v option shows the verbose mod, which means the level of complexity. \
                 For this level, 8 is the maximum value.

    **Command**: *In vim, press i.*

    **Example**: Environment="KUBELET_CONFIG_ARGS=--config=/var/lib/kubelet/config.yaml --hostname-override=ixia-k8s --log-file=/tmp/kubelet.log -v=8"

4.  **Explanation**: Save the changes you have made and exit. \
    **Command**: *In vim, press Esc and :wq.*

5.  **Explanation**: Reload daemon. \
    **Command**: *systemctl daemon-reload*

6.  **Explanation**: Restart kubelet to update the changes. \
    **Command**: *systemctl restart kubelet*

7.  **Explanation**: Check the kubelet status to see if it's in active mod. \
    **Command**: *systemctl status kubelet.*

8.  **Explanation**: To see the messages between the kubernetes components, press the command below. \
    **Command**: *journalctl -u kubelet*


# Debugg kube-apiserver

Prepare enviroment
-

1. Move kube-apiserver.yaml for debugging:\
/etc/kubernetes/manifest/kube-apiserver.yaml -> /etc/kubernetes/kube-apiserver.yaml
2. In /kubernetes/.vscode directory create the folowing files
    * launch.json
    ```json
    {
        "version": "0.2.0",
        "configurations": [
            {
                "name": "Attach remote process",
                "type": "go",
                "request": "attach",
                "mode": "remote",
                "port": 2345,
                "host": "127.0.0.1",
                "apiVersion": 2,
                "showLog": true
            }
        ],
    }
    ```

    * settings.json
    ```json
    {
        "go.alternateTools": {
            "dlv": "${workspaceFolder}/.vscode/dlv-sudo.sh"
        }
    }
    ```

    * dlv-sudo.sh
    ```sh
    #!/bin/sh
    if ! which dlv ; then
        PATH="${GOPATH}/bin:$PATH"
    fi
    if [ "$DEBUG_AS_ROOT" = "true" ]; then
        DLV=$(which dlv)
        exec sudo "$DLV" --only-same-user=false "$@"
    else
        exec dlv "$@"
    fi
    ```
3. 3) run the kube-apiserver process\
$ ./kubernetes/_output/dockerized/bin/linux/amd64/kube-apiserver --advertise-address=10.0.2.15 --allow-privileged=true --authorization-mode=Node,RBAC --client-ca-file=/etc/kubernetes/pki/ca.crt --enable-admission-plugins=NodeRestriction --enable-bootstrap-token-auth=true --etcd-cafile=/etc/kubernetes/pki/etcd/ca.crt --etcd-certfile=/etc/kubernetes/pki/apiserver-etcd-client.crt --etcd-keyfile=/etc/kubernetes/pki/apiserver-etcd-client.key --etcd-servers=https://127.0.0.1:2379 --kubelet-client-certificate=/etc/kubernetes/pki/apiserver-kubelet-client.crt --kubelet-client-key=/etc/kubernetes/pki/apiserver-kubelet-client.key --kubelet-preferred-address-types=InternalIP,ExternalIP,Hostname --proxy-client-cert-file=/etc/kubernetes/pki/front-proxy-client.crt --proxy-client-key-file=/etc/kubernetes/pki/front-proxy-client.key --requestheader-allowed-names=front-proxy-client --requestheader-client-ca-file=/etc/kubernetes/pki/front-proxy-ca.crt --requestheader-extra-headers-prefix=X-Remote-Extra- --requestheader-group-headers=X-Remote-Group --requestheader-username-headers=X-Remote-User --secure-port=6443 --service-account-issuer=https://kubernetes.default.svc.cluster.local --service-account-key-file=/etc/kubernetes/pki/sa.pub --service-account-signing-key-file=/etc/kubernetes/pki/sa.key --service-cluster-ip-range=10.96.0.0/12 --tls-cert-file=/etc/kubernetes/pki/apiserver.crt --tls-private-key-file=/etc/kubernetes/pki/apiserver.key -v=8

4. Attach to the process:\
    * Get the PID\
    $ ps -ef | grep -v grep | grep kube-apiserver | awk '{print $2}' | sed -n '2 p'

    * AS ROOT: $ sudo dlv attach PID ./docker/kube-apiserver --headless --listen=0.0.0.0:2345 --log --api-version 2

    * Run the VSCode debugger

Atention !
-
When we move the file: kube-apiserver.yaml
The clusterul won't work anymore, kubectl get nodes, kubectl get pods\
We will need to move the kube-apiserver.yaml file back at its place.