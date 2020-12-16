#!/bin/bash
docker pull node:10.15-alpine

#Yarn install
docker run -i --rm --name npm_install -u 1000:1000 --tmpfs /home/node/.cache:rw -v $(pwd):/tmp --workdir /tmp node:10.15-alpine npm --verbose install

while [ $(docker ps --filter  name=npm_install* -q | wc -l) -ne 0 ]
do 
	sleep 2
	echo "waiting docker build...."
done 
