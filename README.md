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

# Solve problem with DNS in linux for a certain site

On my virtual machine, I have two network interfaces and for some reason, the internet worked on my operating system, but on my virtual machine, the problem was that if I type ping 8.8.8.8 it works, but if I type ping goole.com, it doesn't work, so the DNS is not working corectly. I found a solutions which could work temporary. Let's solve the problem for a certain site.

Steps
-

1. From /home/ixia, type **sudo vi /etc/hosts**
2. You will see in the first lines some IPV4 adresses. Add one line with the IPV4 address fro the site that you want, the press the space key and type the name of the site. In my case, I typed: **156.140.12.140  bitbucket.it.keysight.com**
3. Press *esc* and *wq* to save changes and quit.
4. Press **ping <name_of_site_given_in_the_file>** to check if it works. I typed **ping bitbucket.it.keysight.com**


