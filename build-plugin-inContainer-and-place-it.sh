#!/bin/bash

echo 'in a new terminal do:'
echo 'step1: ~/kubernetes/build/shell.sh'
echo 'step2: cd ~/go-plugins/demo1-plugins'
echo 'step3: make'
while true; do
read -p "have you done it? [y/n]" yn
    case $yn in
        [Yy]* )
	       ~/kubernetes/build/copy-output.sh;
       		sleep 2;
 		echo 'done copying file from docker to local'		
		sudo cp ~/kubernetes/_output/dockerized/bin/linux/amd64/plugin1.so /usr/bin/plugin1.so; 
		echo 'moving file to /usr/bin/';
		break;;
        [Nn]* ) exit;;
        * ) echo "Please answer yes or no.";;
    esac
done

