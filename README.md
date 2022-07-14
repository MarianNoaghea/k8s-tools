# Documentation

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